package constant

type UserStatusEnum int32

const (
	UserStatusEnum_Active UserStatusEnum = iota + 1
	UserStatusEnum_BlackList
)

func (e UserStatusEnum) Int32() int32 {
	return int32(e)
}

func (e UserStatusEnum) Check() UserStatusEnum {
	switch e {
	case UserStatusEnum_Active:
		return UserStatusEnum_Active
	case UserStatusEnum_BlackList:
		return UserStatusEnum_BlackList
	// undefined
	default:
		return 0
	}
}
