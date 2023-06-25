package reqbody

type Login struct {
	VdID     int32  `json:"vdID" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Vip      int32  `json:"vip" binding:"required"`
	Avatar   int32  `json:"avatar" binding:"required"`
	Language string `json:"language" binding:"required"`
}
