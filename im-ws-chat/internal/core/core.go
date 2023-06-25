package core

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/anypb"

	"imws/internal/cache"
	"imws/internal/constant/exchange"
	"imws/internal/constant/topic"
	"imws/internal/dao"
	"imws/internal/model/entity"
	"imws/internal/model/errcode"
	"imws/internal/model/im"
	"imws/internal/model/vo"
	"imws/internal/util/esutil"
	"imws/internal/util/imredis"
	"imws/internal/util/snowflake"
)

const (
	OK string = "ok"
)

var (
	AVG_MSG_ID_DIFF     int
	CHAT_MSG_CACHE_KEEP int
	MESSAGE_LENGTH_MAX  int
	handler             *commandHandler = &commandHandler{}
)

func init() {
	var err error
	if AVG_MSG_ID_DIFF, err = strconv.Atoi(os.Getenv("AVG_MSG_ID_DIFF")); err != nil {
		log.Fatalf(err.Error())
	}
	if AVG_MSG_ID_DIFF == 0 {
		log.Fatalf("AVG_MSG_ID_DIFF can not be 0")
	}
	if CHAT_MSG_CACHE_KEEP, err = strconv.Atoi(os.Getenv("CHAT_MSG_CACHE_KEEP")); err != nil {
		log.Fatalf(err.Error())
	}
	if AVG_MSG_ID_DIFF == 0 {
		log.Fatalf("CHAT_MSG_CACHE_KEEP can not be 0")
	}
	if MESSAGE_LENGTH_MAX, err = strconv.Atoi(os.Getenv("MESSAGE_LENGTH_MAX")); err != nil {
		log.Fatalf(err.Error())
	}
	if MESSAGE_LENGTH_MAX == 0 {
		log.Fatalf("MESSAGE_LENGTH_MAX can not be 0")
	}
	log.Println("package init ok: core")
}

type userClient struct {
	clients map[*Client]bool
}

type chatRoom struct {
	clients map[*Client]bool
}

type userGroup struct {
	clients map[*Client]bool
}

type subscribeChatVO struct {
	client  *Client
	chatIDs []int32
	reqID   string
}

type commandHandler struct {
}

func (h *commandHandler) validChatMsg(reqID string, msg *im.MessageEntity, c *Client) *im.Push {
	msgInterval := int64(cache.VendorMap[c.info.VdID].MessageInterval)
	now := time.Now()
	ts := now.Unix()
	if ts-c.info.LastMsgTs < msgInterval {
		return &im.Push{
			ReqID:   reqID,
			Command: im.Command_SEND_MESSAGE,
			Code:    0,
			Msg:     "too many message",
		}
	}
	c.info.LastMsgTs = ts
	if utf8.RuneCountInString(msg.Content) > MESSAGE_LENGTH_MAX && msg.ContentType == im.MessageEntity_CHAT {
		return &im.Push{
			ReqID:   reqID,
			Command: im.Command_SEND_MESSAGE,
			Code:    0,
			Msg:     "too much content",
		}
	}

	return nil
}

func (h *commandHandler) validTextingMsg(reqID string, msg *im.TextingMessage, c *Client) *im.Push {
	push := &im.Push{
		ReqID:   reqID,
		Command: im.Command_SEND_TEXTING_MESSAGE,
		Code:    0,
	}
	if msg.GroupID <= 0 {
		push.Msg = "invalid GroupID"
		return push
	}

	return nil
}

func (h *commandHandler) replaceSensitiveWord(userInfo vo.LoginUser, content string) string {
	vendor, isExist := cache.VendorMap[userInfo.VdID]
	if !isExist {
		return content
	}
	words, isExist := vendor.LangSensitiveWord[vendor.ID]
	if !isExist {
		return content
	}
	var newContent string = content
	for sensitiveWord, replacement := range words {
		newContent = strings.Replace(newContent, sensitiveWord, replacement, -1)
	}
	return newContent
}

func cmdHandler(ctx context.Context, req *im.Request, c *Client) {
	log.Printf("cmdHandler command: %s", req.Command.String())
	defer func() {
		r := recover()
		if r != nil {
			log.Println("stacktrace from panic: \n" + string(debug.Stack()))
			push := &im.Push{
				ReqID:   req.ReqID,
				Command: req.Command,
			}
			switch x := r.(type) {
			case string:
				push.Code = 500
				push.Msg = "internal server error: " + x
			case *errcode.Error:
				push.Code = x.Code
				push.Msg = "internal server error: " + x.Detail
			default:
				push.Code = 500
				push.Msg = "internal server error"
			}
			c.onPush <- push
		}
	}()

	switch command := req.Command; command {
	case im.Command_SEND_MESSAGE:
		handler.sendChatMessage(ctx, req, c)
	case im.Command_FETCH_MESSAGES:
		handler.fetchChatMessage(ctx, req, c)
	case im.Command_SUBSCRIBE_CHAT:
		handler.subscribeChat(ctx, req, c)
	case im.Command_REPORT_ABUSE:
		handler.reportAbuse(ctx, req, c)
	default:
		log.Printf("default")
	}
}

func (h *commandHandler) panicIfError(err error) {
	if err != nil {
		panic(errcode.InternalError(err))
	}
}

func (h *commandHandler) sendChatMessage(ctx context.Context, req *im.Request, c *Client) {
	userInfo := c.info
	inBoundMsg := &im.MessageEntity{}
	err := req.Data.UnmarshalTo(inBoundMsg)
	h.panicIfError(err)

	if push := h.validChatMsg(req.ReqID, inBoundMsg, c); push != nil {
		c.onPush <- push
		return
	}

	newContent := h.replaceSensitiveWord(*c.info, inBoundMsg.Content)
	msgId := snowflake.GenerateInt64()

	var visible im.MessageEntity_Visible
	if userInfo.Status == 2 {
		visible = im.MessageEntity_SENDER
	} else {
		visible = im.MessageEntity_ALL
	}

	now := time.Now()
	ts := now.Unix()

	outBoundMsg := &im.MessageEntity{
		MsgID:       msgId,
		ContentType: inBoundMsg.ContentType,
		VdID:        userInfo.VdID,
		Sender:      userInfo.ID,
		SenderName:  userInfo.Name,
		ChatID:      inBoundMsg.ChatID,
		Vip:         userInfo.Vip,
		Avatar:      userInfo.Avatar,
		ReplyTo:     inBoundMsg.ReplyTo,
		Content:     newContent,
		Visible:     visible,
		Timestamp:   ts,
	}

	jsonStr, _ := json.Marshal(outBoundMsg)
	if err := imredis.Client.ZAdd(ctx, imredis.KeyOfChat(inBoundMsg.ChatID), &redis.Z{Score: float64(msgId), Member: jsonStr}).Err(); err != nil {
		h.panicIfError(err)
	}

	PbConvertAndSendMQ(outBoundMsg, exchange.AMQ_TOPIC, topic.CHAT_MESSAGE)
}

func (h *commandHandler) sendTextingMessage(ctx context.Context, req *im.Request, c *Client) {
	userInfo := c.info
	inBoundMsg := &im.TextingMessage{}
	err := req.Data.UnmarshalTo(inBoundMsg)
	h.panicIfError(err)

	if push := h.validTextingMsg(req.ReqID, inBoundMsg, c); push != nil {
		c.onPush <- push
		return
	}

	msgId := snowflake.GenerateInt64()

	now := time.Now()
	ts := now.Unix()

	outBoundMsg := &im.TextingMessage{
		MsgID:       msgId,
		ContentType: inBoundMsg.ContentType,
		VdID:        userInfo.VdID,
		Sender:      userInfo.ID,
		SenderName:  userInfo.Name,
		GroupID:     inBoundMsg.GroupID,
		Avatar:      userInfo.Avatar,
		ReplyTo:     inBoundMsg.ReplyTo,
		Content:     inBoundMsg.Content,
		Timestamp:   ts,
	}

	jsonStr, _ := json.Marshal(outBoundMsg)
	if err := imredis.Client.ZAdd(ctx, imredis.KeyOfGroup(inBoundMsg.GroupID), &redis.Z{Score: float64(msgId), Member: jsonStr}).Err(); err != nil {
		h.panicIfError(err)
	}

	PbConvertAndSendMQ(outBoundMsg, exchange.AMQ_TOPIC, topic.TEXTING_MESSAGE)
}

func (h *commandHandler) subscribeChat(ctx context.Context, req *im.Request, c *Client) {
	a := &im.ChatIDsWrapper{}
	err := req.Data.UnmarshalTo(a)
	h.panicIfError(err)
	c.hub.ClientRoomStateChange(req.ReqID, c, a.GetChatIDs(), true)
}

func (h *commandHandler) fetchChatMessage(ctx context.Context, req *im.Request, c *Client) {
	fetchArgs := &im.FetchArgs{}
	err := req.Data.UnmarshalTo(fetchArgs)
	h.panicIfError(err)
	log.Println(req.Data.String())
	log.Println(fetchArgs.GetPointer())
	log.Println(fetchArgs.Pointer)

	push := &im.Push{
		ReqID:   req.ReqID,
		Command: im.Command_FETCH_MESSAGES,
		Code:    1,
		Msg:     OK,
	}

	var zs []redis.Z
	if fetchArgs.GetPointer() == 0 {
		zs, err = imredis.Client.ZRevRangeWithScores(ctx, imredis.KeyOfChat(fetchArgs.ChatID), 0, 99).Result()
	} else {
		d := &redis.ZRangeBy{
			Min:    strconv.Itoa(0),
			Max:    strconv.FormatInt(fetchArgs.Pointer, 10),
			Offset: 0,
			Count:  99,
		}
		zs, err = imredis.Client.ZRangeByScoreWithScores(ctx, imredis.KeyOfChat(fetchArgs.ChatID), d).Result()
	}

	h.panicIfError(err)

	if len(zs) > 0 { //cache hit
		msgSlice := make([]*im.MessageEntity, len(zs), len(zs))
		for i, z := range zs {
			var msgEntity = &im.MessageEntity{}
			s := z.Member.(string)
			log.Printf("z.Member = %s", s)
			err := json.Unmarshal([]byte(s), msgEntity)
			h.panicIfError(err)
			msgSlice[i] = msgEntity
		}
		data, _ := anypb.New(&im.MessageEntityWrapper{MessageEntity: msgSlice})
		push.Data = data
		c.onPush <- push
		return
	}
	// cache miss, look up elastic search
	list := esutil.GetChatMsg(fetchArgs.ChatID, fetchArgs.Pointer)
	var data *anypb.Any
	if len(list) > 0 {
		data, err = anypb.New(&im.MessageEntityWrapper{MessageEntity: list})
		h.panicIfError(err)
	} else {
		data, _ = anypb.New(&im.MessageEntityWrapper{MessageEntity: make([]*im.MessageEntity, 0)})
	}
	push.Data = data
	c.onPush <- push
}

func (h *commandHandler) fetchTextingMessage(ctx context.Context, req *im.Request, c *Client) {
	arg := &im.FetchTextingMessageArg{}
	err := req.Data.UnmarshalTo(arg)
	h.panicIfError(err)

	// get user group
	userInfo := c.info
	userTimelines := dao.UserTimeline.GetSelf(ctx, &entity.GetSelfTimelineCond{UserId: userInfo.ID, GroupId: arg.GetGroupID(), Pointer: arg.GetPointer()})
	msgIds := make([]int64, 0, len(userTimelines))
	for i := 0; i < len(userTimelines); i++ {
		msgIds = append(msgIds, userTimelines[i].MsgId)
	}
	data, _ := anypb.New(&im.TextingMessageWrapper{Msgs: make([]*im.TextingMessage, 0)})
	// no message
	if len(msgIds) > 0 {
		// get msg by es
		list := esutil.GetTextingMsg(msgIds)
		if len(list) > 0 {
			data, err = anypb.New(&im.TextingMessageWrapper{Msgs: list})
			h.panicIfError(err)
		}
	}

	push := &im.Push{
		ReqID:   req.ReqID,
		Command: im.Command_FETCH_TEXTING_MESSAGES,
		Code:    1,
		Msg:     OK,
		Data:    data,
	}

	c.onPush <- push
}

func (h *commandHandler) reportAbuse(ctx context.Context, req *im.Request, c *Client) {
	args := &im.ReportAbuseArgs{}
	err := req.Data.UnmarshalTo(args)
	h.panicIfError(err)

	tx := dao.GetTx(ctx)

	u := dao.User.GetUser(tx, args.UserID)

	va := &entity.AbuseRecord{
		UserID:   args.GetUserID(),
		ReportBy: c.info.ID,
		Reason:   args.GetReason(),
	}
	dao.User.InsertAbuseRecord(tx, va)

	vd := cache.VendorMap[u.VdID]
	if u.Report >= vd.ReportLock && u.Status == 1 {
		dao.User.IncrReportAndBan(tx, args.GetUserID())
		data, _ := anypb.New(&im.UserStatus{
			Name:   u.Name,
			VdID:   u.VdID,
			Status: 2,
		})
		notify := &im.Notify{
			Command: im.Command_USER_STATUS,
			Data:    data,
		}
		ConvertAndSendMQ(notify, exchange.AMQ_TOPIC, topic.WS_GENERIC)
	} else {
		dao.User.IncrReport(tx, args.GetUserID())
	}

	tx.Commit()

	c.onPush <- &im.Push{
		ReqID:   req.ReqID,
		Code:    1,
		Msg:     OK,
		Command: im.Command_REPORT_ABUSE,
	}
}

func (h *commandHandler) getUserGroups(ctx context.Context, req *im.Request, c *Client) {
	userInfo := c.info
	daoGroups := dao.Group.GetGroups(ctx, userInfo.ID)
	groups := make([]*im.Group, 0, len(daoGroups))
	for i := 0; i < len(daoGroups); i++ {
		groups = append(groups, &im.Group{
			Id:          daoGroups[i].ID,
			Name:        daoGroups[i].Name,
			IconUrl:     daoGroups[i].IconUrl,
			MemberCount: daoGroups[i].MemberCount,
			Deleted:     daoGroups[i].Deleted,
		})
	}

	data, err := anypb.New(&im.GroupsWrapper{Groups: groups})
	h.panicIfError(err)

	c.onPush <- &im.Push{
		ReqID:   req.ReqID,
		Code:    1,
		Msg:     OK,
		Command: im.Command_USER_GROUP,
		Data:    data,
	}
}

func (h *commandHandler) updateReadMessage(ctx context.Context, req *im.Request, c *Client) {
	arg := &im.ReadTextingMessageArg{}
	err := req.Data.UnmarshalTo(arg)
	h.panicIfError(err)
	arg.UserID = c.info.ID

	PbConvertAndSendMQ(arg, exchange.AMQ_TOPIC, topic.READ_MESSAGE)
	c.onPush <- &im.Push{
		ReqID:   req.ReqID,
		Code:    1,
		Msg:     OK,
		Command: im.Command_READ_TEXTING_MESSAGE,
	}
}
