package entity

type GetChatMsgArgs struct {
	ChatID  int64
	Pointer int64
}

type ChatSearchEntity struct {
	Hits *ChatSearchHitsEntity `json:"hits"`
}

type ChatSearchHitsEntity struct {
	Hits []*ChatMessageListEntity `json:"hits"`
}

type ChatMessageListEntity struct {
	Source *MessageEntity `json:"_source"`
}

type MessageEntity struct {
	MsgID       int64  `json:"MsgId"`
	ContentType int32  `json:"ContentType,omitempty"`
	VdID        int64  `json:"VdID,omitempty"`
	Sender      int64  `json:"Sender,omitempty"`
	SenderName  string `json:"SenderName,omitempty"`
	ChatID      int64  `json:"ChatID,omitempty"`
	Vip         int32  `json:"Vip,omitempty"`
	Avatar      int32  `json:"Avatar,omitempty"`
	ReplyTo     int64  `json:"ReplyTo,omitempty"`
	Content     string `json:"Content,omitempty"`
	Visible     int32  `json:"Visible,omitempty"`
	Timestamp   int64  `json:"Timestamp,omitempty"`
}
