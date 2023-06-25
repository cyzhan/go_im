package api

import (
	"net/http"

	"imgateway/internal/rpc/user"
	"imgateway/internal/utils/grpccli"

	"imgateway/internal/model/response"

	"github.com/gin-gonic/gin"
)

var (
	User = &userHandler{}
)

type userHandler struct {
}

func (ud *userHandler) Login(c *gin.Context) {
	var body user.LoginData
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorOf(400, err.Error()))
		return
	}
	conn := grpccli.GetApiConn()
	client := user.NewUserServiceClient(conn)

	result, err := client.Login(c, &body)
	// log.Printf("result = #%v", result)
	// log.Printf("err = %s", err.Error())
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorOf(500, err.Error()))
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, response.ErrorOf(500, result.Error.Msg))
		return
	}
	c.JSON(http.StatusOK, response.Of(result))
}
