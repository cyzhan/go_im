package cache

import (
	"log"

	"imws/internal/dao"
	"imws/internal/model/entity"
)

var (
	VendorMap = make(map[int64]*entity.Vendor)
)

func init() {
	vendors := dao.Vendor.GetVendors()
	for _, vendor := range vendors {
		VendorMap[vendor.ID] = vendor
		vendor.LangSensitiveWord = make(map[int64]map[string]string)
		list := dao.Vendor.GetSensitiveWords(vendor.ID)
		for _, item := range list {
			pair, isExist := vendor.LangSensitiveWord[item.LangID]
			if !isExist {
				pair = map[string]string{}
				vendor.LangSensitiveWord[item.LangID] = pair
			}
			pair[item.Word] = item.Replacement
			// log.Printf("%s : %s", item.Word, item.Replacement)
		}
	}
	log.Printf("package init ok: cache")
}
