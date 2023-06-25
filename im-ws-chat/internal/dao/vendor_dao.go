package dao

import (
	"imws/internal/model/entity"
)

type vendorDao struct {
}

func (vd *vendorDao) GetVendors() []*entity.Vendor {
	s := `SELECT v.id, v.name, v.message_interval, v.message_permission, v.report_lock FROM im.vendor v`
	rows, err := maria.Query(s)
	panicIfError(err)

	var list []*entity.Vendor
	for rows.Next() {
		v := &entity.Vendor{}
		err := rows.Scan(&v.ID, &v.Name, &v.MessageInterval, &v.MessagePermission, &v.ReportLock)
		panicIfError(err)
		list = append(list, v)
	}
	return list
}

func (vd *vendorDao) GetSensitiveWords(vdID int64) []*entity.SensitiveWord {
	s := `SELECT id, vendor_id, lang_id, word, replacement FROM im.sensitive_word WHERE vendor_id = ?`
	rows, err := maria.Query(s, vdID)
	panicIfError(err)

	var list []*entity.SensitiveWord
	for rows.Next() {
		v := &entity.SensitiveWord{}
		err := rows.Scan(&v.ID, &v.VdID, &v.LangID, &v.Word, &v.Replacement)
		panicIfError(err)
		list = append(list, v)
	}
	return list
}
