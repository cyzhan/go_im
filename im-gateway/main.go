package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"imgateway/internal/alpha"
	"imgateway/internal/dao"
	"imgateway/internal/handler/admin"
	"imgateway/internal/handler/api"
	"imgateway/internal/utils/auth"
	"imgateway/internal/utils/encrypt"
	"imgateway/internal/utils/grpccli"
	"imgateway/internal/utils/imredis"
	"imgateway/internal/utils/mysqlcli"
	"imgateway/internal/utils/rabbit"
	"imgateway/internal/utils/snowflake"
)

func main() {
	alpha.InitEnv()
	initPlugin()
	setupAppEnv()
	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		if err, ok := err.(error); ok {
			log.Printf(err.Error())
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "internal server error",
		})
	}))

	server.Use(auth.CORSMiddleware())
	defaultPath := server.Group("/im/gateway")

	apiUsers := defaultPath.Group("/api/users")
	{
		apiUsers.POST("/login", api.User.Login)
	}

	adminUsers := defaultPath.Group("/admin/users")
	{
		adminUsers.Use(auth.PlatformAuth)
		adminUsers.GET("", admin.User.GetList)
		adminUsers.PATCH("", admin.User.Patch)
	}
	adminVendors := defaultPath.Group("/admin/vendors")
	{
		adminVendors.Use(auth.PlatformAuth)
		adminVendors.GET("/sensitive-word", admin.Vendor.GetSensitiveWord)
		adminVendors.PUT("/sensitive-word/:id", admin.Vendor.UpdateSensitiveWord)
		adminVendors.GET("/config", admin.Vendor.GetConfig)
	}
	sys := defaultPath.Group("/system")
	{
		sys.GET("/version", admin.System.GetVersion)
	}

	// By default it serves on :8080 unless a PORT environment variable was defined.
	server.Run(os.Getenv("SERVER_PORT"))
}

func setupAppEnv() {
	value := os.Getenv("ALLOW_ORIGINS")
	origins := strings.Split(value, ",")
	for _, origin := range origins {
		auth.AllowOriginMap[origin] = true
	}

	encrypt.SALT1 = os.Getenv("SALT1")
	auth.JwtSecret = os.Getenv("JWT_SECRET")
}

func initPlugin() {
	snowflake.NewIDGenerator()
	mysqlcli.NewDBClient()
	// minioutil.InitMinio()
	rabbit.InitRabbit()
	dao.InitMaria()
	imredis.InitRedis()
	grpccli.NewGrpc()
}
