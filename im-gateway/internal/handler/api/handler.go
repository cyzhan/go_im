package api

import (
	"imgateway/internal/constant"
	"imgateway/internal/model/bo"

	"github.com/gin-gonic/gin"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func getJwtPayload(ctx *gin.Context) (*bo.JwtPayload, bool) {
	info, ok := ctx.Get(constant.JWT_PAYLOAD)
	if !ok {
		return nil, false
	}

	val := info.(*bo.JwtPayload)
	if val == nil {
		return nil, false
	}

	return val, true
}
