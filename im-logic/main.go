package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"imlogic/internal"
	"imlogic/internal/alpha"
	_ "imlogic/internal/alpha"
	"imlogic/internal/dao"
	"imlogic/internal/service"
	"imlogic/internal/utils/encrypt"
	"imlogic/internal/utils/esutil"
	"imlogic/internal/utils/imredis"
	"imlogic/internal/utils/localcache"
	"imlogic/internal/utils/mysqlcli"
	"imlogic/internal/utils/rabbit"
	"imlogic/internal/utils/snowflake"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	initPlugin()

	rpcServer := internal.NewRpcServer()

	systemCheck()
	rpcServer.Run()

}

func systemCheck() {
	addr := flag.String("addr", os.Getenv("SERVER_PORT"), "http service address")
	http.HandleFunc("/im/logic/system/version", func(w http.ResponseWriter, r *http.Request) {
		service.System.GetVersion(w, r)
	})

	http.Handle("/im/logic/system/metrics", promhttp.Handler())

	server := &http.Server{Addr: *addr, Handler: nil}

	go func() {
		log.Printf("http server start, listen port %s", *addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func initPlugin() {
	alpha.InitEnv()
	snowflake.NewIDGenerator()
	mysqlcli.NewDBClient()
	imredis.InitRedis()
	dao.InitOldMaria()
	localcache.NewCache()
	rabbit.InitRabbit()
	esutil.InitES()

	encrypt.SALT1 = os.Getenv("SALT1")
}
