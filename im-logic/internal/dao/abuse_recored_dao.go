package dao

import (
	"context"

	"gorm.io/gorm"
	"imlogic/internal/model/entity"
)

type abuseRecordDao struct {
}

func NewAbuseRecordDao() *abuseRecordDao {
	return &abuseRecordDao{}
}

func (dao *abuseRecordDao) InsertAbuseRecord(ctx context.Context, db *gorm.DB, args *entity.AbuseRecordEntity) error {
	if err := db.Create(args).Error; err != nil {
		return err
	}

	return nil
}
