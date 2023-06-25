package dao

import "imgateway/internal/model/entity"

type vendorDao struct {
	tokenVdIDMap map[string]*entity.VendorEntity
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

func (vd *vendorDao) GetVdID(token string) int64 {
	return vd.tokenVdIDMap[token].ID
}
