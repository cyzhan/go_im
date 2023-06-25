package core

import (
	"context"
	"log"

	"imws/internal/constant/topic"
	"imws/internal/model/im"
	"imws/internal/rpc/user"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Hub maintains the set of active userClients and broadcasts messages to the
// userClients.
type Hub struct {
	// Registered userClients.
	userClients map[int64]*userClient

	chatRooms map[int64]*chatRoom

	userGroup map[int64]*userGroup

	clientRoomStateChan chan *clientRoomStateDTO

	userGroupsStateChan chan *userGroupsStateDTO

	clientStateChan chan *clientStateDTO
}

func NewHub() *Hub {
	return &Hub{
		userClients:         make(map[int64]*userClient),
		chatRooms:           make(map[int64]*chatRoom),
		userGroup:           make(map[int64]*userGroup),
		clientStateChan:     make(chan *clientStateDTO),
		userGroupsStateChan: make(chan *userGroupsStateDTO),
		clientRoomStateChan: make(chan *clientRoomStateDTO),
	}
}

type clientStateDTO struct {
	isRegister bool
	client     *Client
}

type clientRoomStateDTO struct {
	isJoin  bool
	client  *Client
	chatIDs []int64
	reqID   string
}

type userGroupsStateDTO struct {
	isJoin   bool
	push     *im.Push
	groupIDs []int64
	reqID    string
	client   *Client
}

func RunDefaultHub(ctx context.Context) *Hub {
	hub := NewHub()
	startMQ()
	go hub.run(ctx)
	return hub
}

func (h *Hub) ClientStateChange(client *Client, isRegister bool) {
	h.clientStateChan <- &clientStateDTO{isRegister, client}
}

func (h *Hub) ClientRoomStateChange(reqID string, client *Client, chatIDs []int64, isJoin bool) {
	h.clientRoomStateChan <- &clientRoomStateDTO{
		isJoin:  isJoin,
		client:  client,
		chatIDs: chatIDs,
		reqID:   reqID,
	}
}

func (h *Hub) UserGroupStateChange(client *Client, push *im.Push, groupIDs []int64, isJoin bool) {
	h.userGroupsStateChan <- &userGroupsStateDTO{
		push:     push,
		isJoin:   isJoin,
		groupIDs: groupIDs,
		client:   client,
	}
}

func (h *Hub) run(ctx context.Context) {
	chatMsgs, err := channel.Consume(
		randomQueueNames[topic.CHAT_MESSAGE], // queue
		"",                                   // consumer
		true,                                 // auto-ack
		false,                                // exclusive
		false,                                // no-local
		false,                                // no-wait
		nil,                                  // args
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	notifies, err := channel.Consume(
		randomQueueNames[topic.WS_GENERIC], // queue
		"",                                 // consumer
		true,                               // auto-ack
		false,                              // exclusive
		false,                              // no-local
		false,                              // no-wait
		nil,                                // args
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for {
		select {
		case <-ctx.Done():
			closeMessage := websocket.FormatCloseMessage(1001, "server shutdown")
			for _, userClient := range h.userClients {
				for client, _ := range userClient.clients {
					client.onClose <- closeMessage
				}
			}
			return
		case dto := <-h.clientStateChan:
			if !dto.isRegister {
				if userClient, ok := h.userClients[dto.client.info.ID]; ok {
					delete(userClient.clients, dto.client)
				}
			} else {
				userClients, isExist := h.userClients[dto.client.info.ID]
				if !isExist {
					h.userClients[dto.client.info.ID] = &userClient{clients: make(map[*Client]bool)}
					userClients = h.userClients[dto.client.info.ID]
				}

				userClients.clients[dto.client] = true
			}
		case dto := <-h.clientRoomStateChan:
			for _, chatID := range dto.chatIDs {
				room, isExist := h.chatRooms[chatID]
				if !isExist {
					h.chatRooms[chatID] = &chatRoom{clients: make(map[*Client]bool)}
					room = h.chatRooms[chatID]
				}

				if dto.isJoin {
					log.Println("join")
					room.clients[dto.client] = true
					dto.client.myChat[chatID] = true
				} else {
					delete(room.clients, dto.client)
					delete(dto.client.myChat, chatID)
				}
				chatIDSlice := make([]int64, 0, len(dto.client.myChat))

				for chatID, _ := range dto.client.myChat {
					chatIDSlice = append(chatIDSlice, chatID)
				}

				data, err := anypb.New(&im.ChatIDsWrapper{ChatIDs: chatIDSlice})
				handler.panicIfError(err)

				dto.client.onPush <- &im.Push{
					ReqID:   dto.reqID,
					Command: im.Command_SUBSCRIBE_CHAT,
					Code:    1,
					Msg:     OK,
					Data:    data,
				}
			}
		case d := <-chatMsgs:
			messageEntity := &im.MessageEntity{}
			err := proto.Unmarshal(d.Body, messageEntity)
			if err != nil {
				log.Printf(err.Error())
				break
			}

			data, err := anypb.New(messageEntity)
			if err != nil {
				log.Printf(err.Error())
				break
			}

			push := &im.Push{Command: im.Command_PUSH_MESSAGE, Data: data}
			if chatRoom, isExist := h.chatRooms[messageEntity.GetChatID()]; isExist {
				for client, _ := range chatRoom.clients {
					client.onPush <- push
				}
			}
		case n := <-notifies:
			notify := &im.Notify{}
			if err := proto.Unmarshal(n.Body, notify); err != nil {
				log.Printf(err.Error())
				break
			}

			switch notify.Command {
			case im.Command_USER_STATUS:
				userStatus := user.UserStatusChange{}
				if err := anypb.UnmarshalTo(notify.GetData(), &userStatus, proto.UnmarshalOptions{}); err != nil {
					log.Printf(err.Error())
					break
				}
				client, isExist := h.userClients[userStatus.UserID]
				if !isExist {
					log.Printf("user status change can not find user id: %d", userStatus.UserID)
					break
				}
				for c, _ := range client.clients {
					c.info.Status = userStatus.Status
				}
			default:
				break
			}

		}
	}
}
