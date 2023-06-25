package constant

type GroupMemberTicketTypeEnum int64

const (
	GroupMemberTicketTypeEnum_Invitation GroupMemberTicketTypeEnum = iota + 1
)

type GroupMemberTicketStatusEnum int64

const (
	GroupMemberTicketStatusEnum_Wait GroupMemberTicketStatusEnum = iota
	GroupMemberTicketStatusEnum_Success
	GroupMemberTicketStatusEnum_Fail
)
