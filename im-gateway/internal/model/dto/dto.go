package dto

type ListWrapper struct {
	Items      any   `json:"items"`
	TotalCount int32 `json:"totalCount"`
}

type GetUserListArgs struct {
	Status   int32 `json:"status" form:"status"`
	Page     int32 `json:"page" form:"page" binding:"required,gte=1"`
	Size     int32 `json:"size" form:"size" binding:"required,gte=1,lte=200"`
	VendorID int64 `json:"vendorID"`
}

type GetUserListDTO struct {
	ID          int64   `json:"id,omitempty"`
	AccountType int32   `json:"accountType"`
	Name        string  `json:"name,omitempty"`
	Date        *string `json:"date"`
	Report      *int32  `json:"report"`
	Status      int32   `json:"status,omitempty"`
	Remark      *string `json:"remark"`
}

type PatchUser struct {
	ID       int64   `json:"id" binding:"required"`
	VendorID int64   `json:"vendorID"`
	Status   int32   `json:"status" binding:"required"`
	Remark   *string `json:"remark"`
}
