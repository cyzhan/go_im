package admin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	validate = validator.New()
)

func getVdID(c *gin.Context) int64 {
	value, isExist := c.Get("vdID")
	if !isExist {
		// log.Fatalln("VdID is not found from context")
		panic("VdID is not found from context")
	}

	v, ok := value.(int64)
	if !ok {
		panic("VdID is not reflection ")
	}

	return v
}

func abortBadRequest(c *gin.Context, err error) {
	log.Printf(err.Error())
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  "bad request",
	})
}
