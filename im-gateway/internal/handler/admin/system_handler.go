package admin

import (
	"net/http"

	"imgateway/internal/model/response"

	"github.com/gin-gonic/gin"
)

var (
	System = &systemHandler{"v0.1.0"}
)

type systemHandler struct {
	version string
}

func (hd *systemHandler) GetVersion(c *gin.Context) {
	data := map[string]string{
		"version": hd.version,
	}
	c.JSON(http.StatusOK, response.Of(data))
}
