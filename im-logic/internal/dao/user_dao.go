package dao

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"imlogic/internal/constant"
	"imlogic/internal/model/entity"
	"imlogic/internal/rpc/user"
	_ "imlogic/internal/utils/common"
	"imlogic/internal/utils/mysqlcli"

	"gorm.io/gorm"
)

type userDao struct {
}

func (ud *userDao) NewUser(vendorID int64, name string, pw string) int64 {
	u := &entity.UserEntity{
		VendorID: vendorID,
		Name:     name,
		Password: pw,
		Status:   1,
		Date:     sql.NullString{String: "", Valid: false},
		Report:   0,
		Remark:   "",
	}

	err := mysqlcli.GetClient().Session().
		Table("user").
		Select("vendor_id", "name", "password", "status", "date", "report", "remark").
		Create(u).Error
	panicOnErr(err)
	return u.ID
}

func (ud *userDao) GetUserByName(vdID int64, name string) *entity.UserEntity {
	u := &entity.UserEntity{}

	err := mysqlcli.GetClient().Session().
		Table("user").
		Select("id", "vendor_id", "name", "password", "status", "created_time", "updated_time").
		Where("vendor_id = ?", vdID).
		Where("name = ?", name).First(u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panicOnErr(err)
	}
	return u
}

func (ud *userDao) GetUsers(ctx context.Context, params *user.GetUserListArgs) (list []*user.UserDTO, totalCount int32) {
	sqlFragments := []string{`
	SELECT SQL_CALC_FOUND_ROWS u.name, u.status, IFNULL(u.date, ''), u.report, u.remark 
	FROM im.user u WHERE u.vendor_id = ?`}
	args := []any{params.VendorID}

	if params.Status > 0 {
		sqlFragments = append(sqlFragments, "AND u.status = ?")
		args = append(args, params.Status)
	}

	sqlFragments = append(sqlFragments, "ORDER BY u.id ASC LIMIT ?,?")
	args = append(args, (params.Page-1)*params.Size, params.Size)
	sql := strings.Join(sqlFragments, " ")

	list = make([]*user.UserDTO, 0)
	db := mysqlcli.GetClient().Session()
	rows, err := db.Raw(sql, args...).Rows()
	defer rows.Close()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panicOnErr(err)
	}

	var item *user.UserDTO
	for rows.Next() {
		item = &user.UserDTO{}
		rows.Scan(&item.Name, &item.Status, &item.Date, &item.Report, &item.Remark)
		list = append(list, item)
	}

	err = db.Raw("SELECT FOUND_ROWS() totalCount").Scan(&totalCount).Error
	panicOnErr(err)

	return list, totalCount
}

//func (ud *userDao) PatchUser(ctx context.Context, args *user.PatchUserArgs) error {
//	sqlFragments := []string{`
//	UPDATE im.user
//	SET status=?
//	`}
//	params := []any{args.Status}
//
//	if args.Remark != "" {
//		sqlFragments = append(sqlFragments, ", remark=?")
//		params = append(params, strings.TrimSpace(args.Remark))
//	}
//
//	if args.Status == 2 {
//		sqlFragments = append(sqlFragments, ", date=?")
//		params = append(params, common.UtcMinus4())
//	}
//	sqlFragments = append(sqlFragments, "WHERE id = ? AND vendor_id = ?")
//	params = append(params, args.Id, args.VendorID)
//	//sql := strings.Join(sqlFragments, " ")
//	//
//	//result := tx.Exec(sql, params...)
//	//if result.Error != nil {
//	//	return result.Error
//	//}
//	//affectedRow = result.RowsAffected
//	//return nil
//
//	if err != nil {
//		return affectedRow, err
//	}
//
//	return affectedRow, nil
//}

func (ud *userDao) GetLanguageID(ctx context.Context, vdID int64, langCode string) (langID int32) {
	s := `
	SELECT l.id 
	FROM im.vendor_lang vl 
	INNER JOIN im.lang l ON vl.lang_id = l.id 
	WHERE vl.vendor_id = ? AND l.code = ? AND vl.enable = 1
	`
	err := mysqlcli.GetClient().Session().Raw(s, vdID, langCode).First(&langID).Error
	panicOnErr(err)
	log.Printf("find langID = %d", langID)
	return langID
}

func (ud *userDao) GetUser(ctx context.Context, db *gorm.DB, args *entity.GetUserArgs) (*entity.UserEntity, error) {
	var user *entity.UserEntity
	if err := db.Model(&user).Where("id = ?", args.UserId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ud *userDao) PatchUser(ctx context.Context, db *gorm.DB, args *entity.PatchUserArgs) error {
	db = db.Where("`id` = ?", args.UserId)
	db = db.Where("`vendor_id` = ?", args.VendorId)
	updateMap := make(map[string]any, 0)
	if args.Remark != "" {
		updateMap["remark"] = args.Remark
	}

	if args.Status != constant.UserStatusEnum(0) {
		updateMap["status"] = args.Status
	}

	if args.DateStr != "" {
		updateMap["date"] = args.DateStr
	}

	if args.Report != 0 {
		updateMap["report"] = args.Report
	}

	if err := db.Model(entity.UserEntity{}).Updates(updateMap).Error; err != nil {
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("data RowsAffected  == 0")
	}

	return nil
}
