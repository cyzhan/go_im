package entity

type VendorEntity struct {
	ID                int64  `gorm:"id"`
	Name              string `gorm:"name"`
	MessageInterval   int32  `gorm:"message_interval"`
	Token             string `gorm:"token"`
	MessagePermission int32  `gorm:"message_permission"`
	ReportLock        int32  `gorm:"report_lock"`
}
