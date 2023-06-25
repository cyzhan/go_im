package entity

type VendorEntity struct {
	ID                int64  `gorm:"id"`
	Name              string `gorm:"name"`
	MessageInterval   int32  `gorm:"message_interval"`
	Token             string `gorm:"token"`
	MessagePermission int32  `gorm:"message_permission"`
	ReportLock        int32  `gorm:"report_lock"`
}

func (VendorEntity) TableName() string {
	return "vendor"
}

type GetSensitiveWordPageArgs struct {
	VendorId int64
	Language string
	Page     int32
	Size     int32
}

type GetSensitiveWordListArgs struct {
	VendorId   int64
	LanguageId int64
}

type SensitiveWord struct {
	ID          int64  `gorm:"id"`
	VendorId    int    `gorm:"vendor_id"`
	langId      int    `gorm:"lang_id"`
	Word        string `gorm:"word"`
	Replacement string `gorm:"replacement"`
}

func (SensitiveWord) TableName() string {
	return "sensitive_word"
}

type UpdateSensitiveWordArgs struct {
	Id          int64
	VendorId    int
	Word        string
	Replacement string
}

func (a UpdateSensitiveWordArgs) GenUpdateCond() map[string]any {
	return map[string]any{
		"word":        a.Word,
		"replacement": a.Replacement,
	}
}

type FirstVendorArgs struct {
	VdId int64
}

type LangEntity struct {
	ID          int32  `gorm:"id"`
	Name        string `gorm:"name"`
	DisplayName string `gorm:"display_name"`
	Code        string `gorm:"code"`
}

func (LangEntity) TableName() string {
	return "lang"
}
