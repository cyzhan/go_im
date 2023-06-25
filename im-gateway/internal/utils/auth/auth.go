package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"imgateway/internal/model/bo"

	"imgateway/internal/constant"
	"imgateway/internal/dao"
	"imgateway/internal/utils/imredis"
)

var (
	AllowOriginMap = make(map[string]bool)
	JwtSecret      string
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		value, isExist := AllowOriginMap[origin]
		if value && isExist {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func PlatformAuth(c *gin.Context) {
	if v := c.Request.Header["Authorization"]; len(v) > 0 {
		value := v[0]
		token := strings.Replace(value, "Bearer ", "", 1)
		if vdID := dao.Vendor.GetVdID(token); vdID > 0 {
			c.Set("vdID", vdID)
			c.Next()
			return
		}
	}
	abort(c)
	return
}

func SportAuth(c *gin.Context) {
	values := c.Request.Header["Authorization"]
	if len(values) == 0 {
		abort(c)
		return
	}

	tokenString := strings.Replace(values[0], "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(JwtSecret), nil
	})

	if err != nil {
		log.Printf(err.Error())
		abort(c)
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fragments := strings.Split(tokenString, ".")
		signature := fragments[2]
		if err := imredis.Client.Get(c, signature).Err(); err != nil {
			log.Printf(err.Error())
			abort(c)
			return
		}

		if data, err := base64.RawURLEncoding.DecodeString(fragments[1]); err != nil {
			log.Printf(err.Error())
			abort(c)
			return
		} else {
			vo := &bo.JwtPayload{}
			json.Unmarshal(data, vo)
			c.Set(constant.JWT_PAYLOAD, vo)
			c.Next()
			return
		}
	} else {
		fmt.Printf(err.Error())
		abort(c)
		return
	}
}

func abort(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  "unAuthorized",
	})
}
