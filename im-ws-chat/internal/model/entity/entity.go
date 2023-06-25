package entity

import (
	`database/sql`
)

type Vendor struct {
	ID                int64                       `json:"id"`
	Name              string                      `json:"name"`
	MessageInterval   int32                       `json:"messageInterval"`
	MessagePermission int32                       `json:"messagePermission"`
	ReportLock        int32                       `json:"reportLock"`
	LangSensitiveWord map[int64]map[string]string `json:"langSensitiveWord"`
}

type SensitiveWord struct {
	ID          int64
	VdID        int64
	LangID      int64
	Word        string
	Replacement string
}

type User struct {
	ID       int64  `json:"id"`
	VdID     int64  `json:"vdID"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int32  `json:"status"`
	// Date     string `json:"date"`
	Date   sql.NullString `json:"date"`
	Report int32          `json:"report"`
	Remark string         `json:"remark"`
}

type AbuseRecord struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userID"`
	VdID        int64  `json:"vdID"`
	Reason      string `json:"reason"`
	ReportBy    int64  `json:"reportBy"`
	CreatedTime string `json:"createdTime"`
}

type UserGroups struct {
	UserGroups []UserGroup `json:"groups"`
}

type UserGroup struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	IconUrl     string `json:"iconUrl"`
	MemberCount int32  `json:"memberCount"`
	Deleted     int32  `json:"deleted"`
}

type UserTimeline struct {
	MsgId   int64 `json:"msgId"`
	UserId  int64 `json:"userId"`
	GroupId int64 `json:"groupId"`
}

type GetSelfTimelineCond struct {
	UserId  int64
	GroupId int64
	Pointer int64
}
