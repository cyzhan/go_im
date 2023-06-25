package dto

type HelloReq struct {
	Name string `json:"name"`
}

type HelloResp struct {
	Msg string `json:"msg"`
}
