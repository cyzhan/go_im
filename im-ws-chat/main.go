package main

import (
	"context"
	"flag"

	_ "imws/internal/alpha"
	_ "imws/internal/cache"
	"imws/internal/core"
	_ "imws/internal/util/esutil"
	"imws/internal/util/snowflake"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addr = flag.String("addr", os.Getenv("SERVER_PORT"), "http service address")

func main() {
	flag.Parse()
	initPlugin()

	ctx, cancel := context.WithCancel(context.Background())
	hub := core.RunDefaultHub(ctx)
	http.HandleFunc("/im", func(w http.ResponseWriter, r *http.Request) {
		core.ServeWs(ctx, hub, w, r)
	})

	server := &http.Server{Addr: *addr, Handler: nil}

	go func() {
		log.Printf("server start")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	defer func() {
		cancel()
		time.Sleep(2 * time.Second)
		log.Println("Server exiting")
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

}

func initPlugin() {
	snowflake.NewIDGenerator()
}
