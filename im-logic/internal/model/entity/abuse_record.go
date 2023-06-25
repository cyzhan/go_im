package entity

import (
	"time"
)

type AbuseRecordEntity struct {
	Id          int64     `gorm:"primaryKey;autoIncrement;id"`
	UserId      int64     `gorm:"user_id"`
	VendorId    int64     `gorm:"vendor_id"`
	Reason      string    `gorm:"Reason"`
	ReportedBy  int64     `gorm:"reported_by"`
	CreatedTime time.Time `gorm:"created_time"`
	UpdatedTime time.Time `gorm:"updated_time"`
}

func (AbuseRecordEntity) TableName() string {
	return "abuse_record"
}
