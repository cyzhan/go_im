package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"imgateway/internal/model/dto"
	"imgateway/internal/model/response"
	"imgateway/internal/rpc/vendors"
	"imgateway/internal/utils/grpccli"
)

type vendorHandler struct {
}

var Vendor = &vendorHandler{}

func (vd *vendorHandler) UpdateSensitiveWord(c *gin.Context) {
	var args = dto.SensitiveWord{VdID: getVdID(c)}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		abortBadRequest(c, err)
		return
	}

	args.ID = id

	if err := c.BindJSON(&args); err != nil {
		abortBadRequest(c, err)
		return
	}

	rpcSensitiveWord := &vendors.SensitiveWord{
		Id:          args.ID,
		Word:        args.Word,
		Replacement: args.Replacement,
	}
	rpcReq := &vendors.UpdateSensitiveWordReq{
		SensitiveWord: rpcSensitiveWord,
		VdId:          args.VdID,
	}

	conn := grpccli.GetApiConn()
	client := vendors.NewVendorsServiceClient(conn)
	if _, err := client.UpdateSensitiveWord(c, rpcReq); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorOf(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.OK())
}

func (vd *vendorHandler) GetSensitiveWord(c *gin.Context) {
	var args = dto.GetSensitiveWordArgs{VdID: getVdID(c)}
	if err := c.BindQuery(&args); err != nil {
		abortBadRequest(c, err)
		return
	}

	rpcReq := &vendors.GetSensitiveWordPageReq{
		Language:  args.Language,
		PageIndex: args.Page,
		PageSize:  args.Size,
		VdId:      args.VdID,
	}
	conn := grpccli.GetApiConn()
	client := vendors.NewVendorsServiceClient(conn)
	rpcResp, err := client.GetSensitiveWordPage(c, rpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorOf(400, err.Error()))
		return
	}

	wrapper := dto.ListWrapper{
		Items:      rpcResp.SensitiveWords,
		TotalCount: rpcResp.TotalCount,
	}
	c.JSON(http.StatusOK, response.Of(wrapper))
}

func (vd *vendorHandler) GetConfig(c *gin.Context) {
	var args = dto.VendorConfigReq{VdId: getVdID(c)}
	conn := grpccli.GetApiConn()
	rpcSrv := vendors.NewVendorsServiceClient(conn)
	rpcReq := &vendors.GetConfigReq{VdId: args.VdId}
	config, err := rpcSrv.GetConfig(c, rpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorOf(400, err.Error()))
		return
	}

	enableLang := make([]*dto.Lang, 0, len(config.GetEnableLang()))
	for i := 0; i < len(config.GetEnableLang()); i++ {
		enableLang = append(enableLang, &dto.Lang{
			Id:   config.GetEnableLang()[i].Id,
			Name: config.GetEnableLang()[i].Name,
			Code: config.GetEnableLang()[i].Code,
		})
	}
	resp := &dto.VendorConfigResp{
		MessageInterval:   config.MessageInterval,
		MessagePermission: config.MessagePermission,
		ReportLock:        config.ReportLock,
		EnableLan:         enableLang,
	}
	c.JSON(http.StatusOK, response.Of(resp))
}
