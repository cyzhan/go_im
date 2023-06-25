package core

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"imws/internal/dao"
	"imws/internal/model/im"
	"imws/internal/model/vo"
	"imws/internal/util/imredis"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	conn *websocket.Conn

	info *vo.LoginUser

	myChat map[int64]bool

	groups map[int64]bool

	onPush chan *im.Push

	onClose chan []byte
}

type RoomState struct {
	reqID   string
	chatIDs []int32
}

// serveWs handles websocket requests from the peer.
func ServeWs(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tokens, present := query["token"]
	if !present || len(tokens) == 0 {
		log.Println("tokens not present")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	value, err := imredis.Client.Get(context.Background(), tokens[0]).Result()
	if value == "" || err != nil {
		log.Println("invalid token")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var userInfo vo.LoginUser

	if err := json.Unmarshal([]byte(value), &userInfo); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:     hub,
		conn:    conn,
		info:    &userInfo,
		myChat:  make(map[int64]bool),
		groups:  make(map[int64]bool),
		onPush:  make(chan *im.Push, 64),
		onClose: make(chan []byte),
	}

	hub.ClientStateChange(client, true)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump(ctx)
	//load user group info
	// go client.userGroupRegister(ctx)
	go client.pushInitData(ctx)
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump(ctx context.Context) {
	defer func() {
		c.hub.ClientStateChange(c, false)
		close(c.onClose)
		c.conn.Close()
		log.Printf("%s readPump goroutine close", c.info.Name)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf("onReceive msg")
		readCtx, cancel := context.WithCancel(ctx)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			cancel()
			break
		}

		req := &im.Request{}
		if err = proto.Unmarshal(message, req); err != nil {
			log.Printf("proto 解析失敗 %v", err)
			cancel()
			break
		}

		cmdHandler(readCtx, req, c)
		cancel()
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
		log.Printf("%s writePump goroutine close", c.info.Name)
	}()
	for {
		select {
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf(err.Error())
				break
			}
		case push := <-c.onPush:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			log.Printf("ReqID=%s, command:%d", push.ReqID, push.Command)
			bytes, err := proto.Marshal(push)
			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				break
			}
			w.Write(bytes)

			if err := w.Close(); err != nil {
				break
			}
		case closeMessage, ok := <-c.onClose:
			if !ok {
				log.Printf("client side close")
				return
			}
			c.conn.WriteControl(websocket.CloseMessage, closeMessage, time.Now().Add(time.Second))
			return
		}
	}
}

func (c *Client) pushInitData(ctx context.Context) {
	args := &im.FetchTextingMessageArg{
		Pointer: 0,
		GroupID: 0,
	}
	data, err := anypb.New(args)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	req := &im.Request{
		ReqID:   "",
		Command: im.Command_FETCH_MESSAGES,
		Data:    data,
	}

	handler.fetchTextingMessage(ctx, req, c)

	list := dao.UserTimeline.GetUserGroupPointer(ctx, c.info.ID)
	data, err = anypb.New(&im.ReadPointerWrapper{Pointers: list})
	if err != nil {
		log.Printf(err.Error())
		return
	}

	push := &im.Push{
		Command: im.Command_FETCH_READ_POINTERS,
		Code:    1,
		Msg:     OK,
		Data:    data,
	}

	c.onPush <- push

	wrapper, ids := getUserGroups(ctx, c.info.ID)

	data, _ = anypb.New(wrapper)
	push = &im.Push{
		ReqID:   "1",
		Code:    1,
		Msg:     OK,
		Command: im.Command_USER_GROUP,
		Data:    data,
	}

	c.hub.UserGroupStateChange(c, push, ids, true)
}

func (c *Client) userGroupRegister(ctx context.Context) {
	wrapper, ids := getUserGroups(ctx, c.info.ID)

	data, _ := anypb.New(wrapper)
	push := &im.Push{
		ReqID:   "1",
		Code:    1,
		Msg:     OK,
		Command: im.Command_USER_GROUP,
		Data:    data,
	}

	c.hub.UserGroupStateChange(c, push, ids, true)

}
func getUserGroups(ctx context.Context, userId int64) (*im.GroupsWrapper, []int64) {
	daoGroups := dao.Group.GetGroups(ctx, userId)
	groups := make([]*im.Group, 0, len(daoGroups))
	groupIDs := make([]int64, 0, len(daoGroups))
	for i := 0; i < len(daoGroups); i++ {
		// hub register
		groupIDs = append(groupIDs, daoGroups[i].ID)
		// make protobuf response
		groups = append(groups, &im.Group{
			Id:          daoGroups[i].ID,
			Name:        daoGroups[i].Name,
			IconUrl:     daoGroups[i].IconUrl,
			MemberCount: daoGroups[i].MemberCount,
			Deleted:     daoGroups[i].Deleted,
		})
	}

	return &im.GroupsWrapper{Groups: groups}, groupIDs
}

func (c *Client) OnPush(push *im.Push) {
	c.onPush <- push
}
