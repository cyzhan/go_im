package entity

type UserEntity struct {
	ID          int32   `json:"id,omitempty"`
	VendorID    int32   `json:"vendorID,omitempty"`
	AccountType int32   `json:"accountType,omitempty"`
	Name        string  `json:"name,omitempty"`
	Password    string  `json:"password,omitempty"`
	Date        *string `json:"date,omitempty"`
	Report      *int32  `json:"report,omitempty"`
	Vip         int32   `json:"vip,omitempty"`
	Avatar      int32   `json:"avatar,omitempty"`
	Status      int32   `json:"status,omitempty"`
	Remark      *string `json:"remark,omitempty"`
	CreatedTime string  `json:"createdTime,omitempty"`
	UpdatedTime string  `json:"updatedTime,omitempty"`
}
