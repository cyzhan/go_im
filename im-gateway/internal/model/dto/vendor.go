package dto

type GetSensitiveWordArgs struct {
	VdID     int64
	Language string `json:"language" form:"language" binding:"required"`
	Page     int32  `json:"page" form:"page" binding:"required,gte=1"`
	Size     int32  `json:"size" form:"size" binding:"required,gte=1,lte=200"`
}

type SensitiveWord struct {
	ID          int64  `json:"id" gorm:"id"`
	Word        string `json:"word" binding:"required" gorm:"word"`
	Replacement string `json:"replacement" binding:"required" gorm:"replacement"`
	VdID        int64
}

type VendorConfigReq struct {
	VdId int64
}

type VendorConfigResp struct {
	MessageInterval   int32   `json:"messageInterval"`
	MessagePermission int32   `json:"messagePermission"`
	ReportLock        int32   `json:"reportLock"`
	EnableLan         []*Lang `json:"enableLang"`
}

type Lang struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
