package service

import (
	"context"
	"log"

	"imlogic/internal/dao"
	"imlogic/internal/model/entity"
	"imlogic/internal/rpc/common"
	"imlogic/internal/rpc/message"
)

type messageService struct {
}

func NewMessageService() *messageService {
	return &messageService{}
}

func (s *messageService) FetchMessage(ctx context.Context, req *message.FetchMessageReq) (*message.FetchMessageResp, error) {
	// cache miss, look up elastic search
	list, err := dao.ESMessageDao.GetChatMsg(ctx, &entity.GetChatMsgArgs{
		ChatID:  req.ChatID,
		Pointer: req.Pointer,
	})
	if err != nil {
		log.Println("(s *messageService) FetchMessage dao.ESMessageDao.GetChatMsg err : %s", err.Error())
		defaultErr := &common.Error{
			Code: 2,
			Msg:  "data no rows",
		}
		return &message.FetchMessageResp{Error: defaultErr}, nil

	}

	msg := make([]*message.Message, 0, len(list))
	for _, l := range list {
		msg = append(msg, &message.Message{
			MsgID:       l.MsgID,
			ContentType: l.ContentType,
			VdID:        l.VdID,
			Sender:      l.Sender,
			SenderName:  l.SenderName,
			ChatID:      l.ChatID,
			Vip:         l.Vip,
			Avatar:      l.Avatar,
			ReplyTo:     l.ReplyTo,
			Content:     l.Content,
			Visible:     l.Visible,
			Timestamp:   l.Timestamp,
		})
	}

	return &message.FetchMessageResp{
		MessageWrap: msg,
	}, nil
}
