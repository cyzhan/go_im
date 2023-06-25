package response

var ok = &map[string]any{
	"code": 1,
	"msg":  "ok",
}

func OK() *map[string]any {
	return ok
}

func Of[T any](v T) *map[string]any {
	return &map[string]any{
		"code": 1,
		"msg":  "ok",
		"data": v,
	}
}

func ErrorOf(code int32, msg string) *map[string]any {
	return &map[string]any{
		"code": code,
		"msg":  msg,
	}
}
