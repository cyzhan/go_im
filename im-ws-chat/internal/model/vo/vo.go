package vo

type LoginUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	VdID      int64  `json:"vdID"`
	Vip       int32  `json:"vip"`
	Status    int32  `json:"status"`
	Avatar    int32  `json:"avatar"`
	LangID    int64  `json:"langID"`
	LastMsgTs int64
}

type CacheClean struct {
	ChatID int64 `json:"chatID"`
	MsgID  int64 `json:"msgID"`
}

type Group struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	IconUrl     string `json:"iconUrl"`
	MemberCount int32  `json:"memberCount"`
	Deleted     int32  `json:"deleted"`
}
