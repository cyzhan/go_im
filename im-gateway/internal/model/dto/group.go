package dto

type CreateGroupReq struct {
	Name     string  `json:"name" binding:"required"`
	Invitees []int64 `json:"invitees" binding:"required"`
}

type UpdateGroupReq struct {
	Name    string `json:"name" binding:"required"`
	IconUrl string `json:"iconUrl" binding:"required"`
}

type SendInvitationReq struct {
	GroupId  int64   `json:"groupId" binding:"required"`
	Invitees []int64 `json:"invitees" binding:"required"`
}
