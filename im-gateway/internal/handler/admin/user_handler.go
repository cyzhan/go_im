package admin

import (
	"log"
	"net/http"

	"imgateway/internal/model/dto"
	"imgateway/internal/model/response"
	"imgateway/internal/rpc/user"
	"imgateway/internal/utils/grpccli"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

var (
	User = &userHandler{}
)

func (uh *userHandler) GetList(c *gin.Context) {
	var reqArgs dto.GetUserListArgs
	if err := c.BindQuery(&reqArgs); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	grpcArgs := user.GetUserListArgs{
		Status:   reqArgs.Status,
		Page:     reqArgs.Page,
		Size:     reqArgs.Size,
		VendorID: getVdID(c),
	}

	conn := grpccli.GetApiConn()
	client := user.NewUserServiceClient(conn)

	result, err := client.GetUserList(c, &grpcArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorOf(http.StatusBadRequest, err.Error()))
		return
	}
	if result.Error.GetCode() != 0 {
		c.JSON(http.StatusBadRequest, response.ErrorOf(http.StatusBadRequest, result.Error.Msg))
		return
	}

	list := make([]*dto.GetUserListDTO, 0)
	for _, item := range result.Items {
		list = append(list, &dto.GetUserListDTO{
			ID:          item.Id,
			Name:        item.Name,
			AccountType: item.AccountType,
			Date:        &item.Date,
			Report:      &item.Report,
			Status:      item.Status,
			Remark:      &item.Remark,
		})
	}

	wrapper := dto.ListWrapper{
		Items:      list,
		TotalCount: result.TotalCount,
	}
	c.JSON(http.StatusOK, response.Of(wrapper))
}

func (uh *userHandler) Patch(c *gin.Context) {
	var params dto.PatchUser
	if err := c.BindJSON(&params); err != nil {
		log.Printf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "bad request",
		})
		return
	}

	grpcArgs := user.PatchUserArgs{
		VendorID: getVdID(c),
		Id:       params.ID,
		Status:   params.Status,
		Remark:   *params.Remark,
	}

	conn := grpccli.GetApiConn()
	client := user.NewUserServiceClient(conn)

	commonErr, err := client.PatchUser(c, &grpcArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorOf(500, err.Error()))
		return
	}
	if commonErr.GetCode() != 0 {
		c.JSON(http.StatusBadRequest, response.ErrorOf(int32(commonErr.GetCode()), commonErr.Msg))
		return
	}

	c.JSON(http.StatusOK, response.OK())
}
