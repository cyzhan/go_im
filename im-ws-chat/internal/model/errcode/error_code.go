package errcode

type Error struct {
	Code   int32
	Msg    string
	Detail string
}

func (e *Error) Error() string {
	return e.Detail
}

func InternalError(err error) error {
	return &Error{
		Code:   500,
		Msg:    "internal error",
		Detail: err.Error(),
	}
}
