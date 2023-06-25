package service

import (
	"context"
	"strconv"
	"sync"
	"time"

	"imlogic/internal/constant/cacheKey"
	"imlogic/internal/dao"
	"imlogic/internal/model/entity"
	"imlogic/internal/rpc/vendors"
	"imlogic/internal/utils/localcache"
	"imlogic/internal/utils/mysqlcli"
)

func NewVendorsService() *vendorsService {
	return &vendorsService{}
}

type vendorsService struct {
	sensitiveWordListLock sync.Mutex
}

func (s *vendorsService) UpdateSensitiveWord(ctx context.Context, req *vendors.UpdateSensitiveWordReq) (*vendors.UpdateSensitiveWordResp, error) {
	db := mysqlcli.GetClient().Session()
	if err := dao.VendorDao.UpdateSensitiveWord(ctx, db, &entity.UpdateSensitiveWordArgs{
		Id:          req.SensitiveWord.Id,
		VendorId:    int(req.VdId),
		Word:        req.SensitiveWord.Word,
		Replacement: req.SensitiveWord.Replacement,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *vendorsService) GetSensitiveWordPage(ctx context.Context, req *vendors.GetSensitiveWordPageReq) (*vendors.GetSensitiveWordPageResp, error) {
	db := mysqlcli.GetClient().Session()
	list, totalCount, err := dao.VendorDao.GetSensitiveWordPage(ctx, db, &entity.GetSensitiveWordPageArgs{
		Language: req.Language,
		Page:     req.PageIndex,
		Size:     req.PageSize,
		VendorId: req.VdId,
	})
	if err != nil {
		return nil, err
	}

	respList := make([]*vendors.SensitiveWord, 0, len(list))
	for i := 0; i < len(list); i++ {
		respList = append(respList, &vendors.SensitiveWord{
			Id:          list[i].ID,
			Word:        list[i].Word,
			Replacement: list[i].Replacement,
		})
	}

	resp := &vendors.GetSensitiveWordPageResp{
		SensitiveWords: respList,
		TotalCount:     totalCount,
	}

	return resp, nil
}

func (s *vendorsService) GetConfig(ctx context.Context, req *vendors.GetConfigReq) (*vendors.GetConfigResp, error) {
	db := mysqlcli.GetClient().Session()
	vendor, err := dao.VendorDao.FirstVendor(ctx, db, &entity.FirstVendorArgs{VdId: req.GetVdId()})
	if err != nil {
		return nil, err
	}

	langs, err := dao.VendorDao.GetVendorLang(ctx, db, &entity.FirstVendorArgs{VdId: vendor.ID})
	if err != nil {
		return nil, err
	}

	respLangs := make([]*vendors.Lang, 0, len(langs))
	for i := 0; i < len(langs); i++ {
		respLangs = append(respLangs, &vendors.Lang{
			Id:   langs[i].ID,
			Name: langs[i].DisplayName,
			Code: langs[i].Code,
		})
	}

	resp := &vendors.GetConfigResp{
		MessageInterval:   vendor.MessageInterval,
		MessagePermission: vendor.MessagePermission,
		ReportLock:        vendor.ReportLock,
		EnableLang:        respLangs,
	}

	return resp, nil
}

func (s *vendorsService) GetSensitiveWordList(ctx context.Context, req *vendors.GetSensitiveWordListReq) (*vendors.GetSensitiveWordListResp, error) {
	s.sensitiveWordListLock.Lock()
	defer s.sensitiveWordListLock.Unlock()
	cacheKey := s.getSensitiveWordListCacheKey(ctx, req.GetVdId(), req.GetLanguageId())
	if data, ok := localcache.Cache().Get(cacheKey); ok {
		if list, ok := data.([]*vendors.SensitiveWord); ok {
			resp := &vendors.GetSensitiveWordListResp{}
			resp.SensitiveWords = list
			return resp, nil
		}
	}

	db := mysqlcli.GetClient().Session()
	list, err := dao.VendorDao.GetSensitiveWordList(ctx, db, &entity.GetSensitiveWordListArgs{
		LanguageId: req.LanguageId,
		VendorId:   req.VdId,
	})
	if err != nil {
		return nil, err
	}

	respList := make([]*vendors.SensitiveWord, 0, len(list))
	for i := 0; i < len(list); i++ {
		respList = append(respList, &vendors.SensitiveWord{
			Id:          list[i].ID,
			Word:        list[i].Word,
			Replacement: list[i].Replacement,
		})
	}

	resp := &vendors.GetSensitiveWordListResp{
		SensitiveWords: respList,
	}

	localcache.Cache().SaveWithExpiration(cacheKey, respList, 6*time.Hour)
	return resp, nil
}

func (s *vendorsService) getSensitiveWordListCacheKey(ctx context.Context, vendorId, langId int64) string {
	return cacheKey.CacheVendorSensitiveWord + ":" + strconv.FormatInt(vendorId, 10) + ":" + strconv.FormatInt(langId, 10)
}
