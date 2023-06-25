package entity

import (
	"database/sql"

	"imlogic/internal/constant"
)

type UserEntity struct {
	ID          int64          `json:"id,omitempty" gorm:"id"`
	VendorID    int64          `json:"vendorID,omitempty" gorm:"vendor_id"`
	AccountType int32          `json:"accountType,omitempty" gorm:"account_type"`
	Name        string         `json:"name,omitempty"`
	Password    string         `json:"password,omitempty"`
	Date        sql.NullString `json:"date,omitempty"`
	Report      int32          `json:"report"`
	Vip         int32          `json:"vip,omitempty"`
	Avatar      int32          `json:"avatar,omitempty"`
	Status      int32          `json:"status,omitempty"`
	Remark      string         `json:"remark"`
	CreatedTime string         `json:"createdTime,omitempty" gorm:"created_time"`
	UpdatedTime string         `json:"updatedTime,omitempty" gorm:"updated_time"`
}

func (UserEntity) TableName() string {
	return "user"
}

type GetUserArgs struct {
	UserId int64
}

type PatchUserArgs struct {
	Remark  string
	Status  constant.UserStatusEnum
	DateStr string
	Report  int32
	// where
	UserId   int64
	VendorId int64
}
