package dao

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
	"imlogic/internal/constant/cacheKey"
	"imlogic/internal/model/entity"
	"imlogic/internal/utils/localcache"
)

type vendorDao struct {
	sync.Mutex
	tokenVdIDMapCacheKey string
	vdIDMapCacheKey      string
}

func NewVendorDao() *vendorDao {
	return &vendorDao{tokenVdIDMapCacheKey: cacheKey.CacheTokenVendorIdMap, vdIDMapCacheKey: cacheKey.CacheVendorIdMap}
}

func (vd *vendorDao) GetVendors() []*entity.VendorEntity {
	s := `SELECT v.id, v.name, v.token, v.message_interval, v.message_permission, v.report_lock FROM im.vendor v`
	rows, err := maria.Query(s)
	panicOnErr(err)

	var list []*entity.VendorEntity
	for rows.Next() {
		v := &entity.VendorEntity{}
		err := rows.Scan(&v.ID, &v.Name, &v.Token, &v.MessageInterval, &v.MessagePermission, &v.ReportLock)
		panicOnErr(err)
		list = append(list, v)
	}
	return list
}

func (vd *vendorDao) GetSensitiveWordPage(ctx context.Context, db *gorm.DB, args *entity.GetSensitiveWordPageArgs) ([]*entity.SensitiveWord, int32, error) {
	words := make([]*entity.SensitiveWord, 0, 0)

	if err := db.Model(entity.SensitiveWord{}).Joins("INNER JOIN im.lang AS la ON sensitive_word.lang_id = la.id").
		Where("sensitive_word.vendor_id = ?", args.VendorId).Where("la.code = ?", args.Language).
		Offset(int(offset(args.Page, args.Size))).Limit(int(args.Size)).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	var totalCount int64
	if err := db.Model(entity.SensitiveWord{}).Joins("INNER JOIN im.lang AS la ON sensitive_word.lang_id = la.id").
		Where("sensitive_word.vendor_id = ?", args.VendorId).Where("la.code = ?", args.Language).
		Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return words, int32(totalCount), nil
}

func (vd *vendorDao) UpdateSensitiveWord(ctx context.Context, db *gorm.DB, args *entity.UpdateSensitiveWordArgs) error {
	db = db.Where("`id` = ?", args.Id)
	db = db.Where("`vendor_id` = ?", args.VendorId)

	if err := db.Model(entity.SensitiveWord{}).Updates(args.GenUpdateCond()).Error; err != nil {
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("data RowsAffected no row")
	}

	return nil
}

func (vd *vendorDao) GetTokenVdMap() (map[string]*entity.VendorEntity, error) {
	vd.Mutex.Lock()
	defer vd.Mutex.Unlock()

	if data, ok := localcache.Cache().Get(vd.tokenVdIDMapCacheKey); ok {
		vendor, ok := data.(map[string]*entity.VendorEntity)
		if ok {
			return vendor, nil
		}
	}

	vendorList := vd.GetVendors()
	tokenVdIDMap := make(map[string]*entity.VendorEntity, len(vendorList))
	for i := 0; i < len(vendorList); i++ {
		tokenVdIDMap[vendorList[i].Token] = vendorList[i]
	}

	//  1209600*time.Second = > 2week
	localcache.Cache().SaveWithExpiration(vd.tokenVdIDMapCacheKey, tokenVdIDMap, 1209600*time.Second)
	return tokenVdIDMap, nil
}

func (vd *vendorDao) GetVdIDMap() (map[int64]*entity.VendorEntity, error) {
	vd.Mutex.Lock()
	defer vd.Mutex.Unlock()

	if data, ok := localcache.Cache().Get(vd.vdIDMapCacheKey); ok {
		vendor, ok := data.(map[int64]*entity.VendorEntity)
		if ok {
			return vendor, nil
		}
	}

	vendorList := vd.GetVendors()
	vdIDMap := make(map[int64]*entity.VendorEntity, len(vendorList))
	for i := 0; i < len(vendorList); i++ {
		vdIDMap[vendorList[i].ID] = vendorList[i]
	}

	//  1209600*time.Second = > 2week
	localcache.Cache().SaveWithExpiration(vd.vdIDMapCacheKey, vdIDMap, 1209600*time.Second)
	return vdIDMap, nil
}

func (vd *vendorDao) FirstVendor(ctx context.Context, db *gorm.DB, args *entity.FirstVendorArgs) (*entity.VendorEntity, error) {
	var vendor *entity.VendorEntity
	if err := db.Model(&vendor).Where("id = ?", args.VdId).First(&vendor).Error; err != nil {
		return nil, err
	}
	fmt.Println(vendor)

	return vendor, nil
}

func (vd *vendorDao) GetVendorLang(ctx context.Context, db *gorm.DB, args *entity.FirstVendorArgs) ([]*entity.LangEntity, error) {
	langs := make([]*entity.LangEntity, 0)
	if err := db.Model(entity.LangEntity{}).Joins("INNER JOIN vendor_lang AS vl ON vl.lang_id = lang.id").
		Where("vl.vendor_id = ?", args.VdId).Find(&langs).Error; err != nil {
		return nil, err
	}

	return langs, nil
}

func (vd *vendorDao) GetSensitiveWordList(ctx context.Context, db *gorm.DB, args *entity.GetSensitiveWordListArgs) ([]*entity.SensitiveWord, error) {
	words := make([]*entity.SensitiveWord, 0, 0)

	if err := db.Model(entity.SensitiveWord{}).
		Where("vendor_id = ?", args.VendorId).Where("lang_id = ?", args.LanguageId).
		Find(&words).Error; err != nil {
		return nil, err
	}

	return words, nil
}
