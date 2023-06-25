package imredis

import (
	"context"
	"log"
	"os"

	"strconv"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis() {
	poolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	panicOnErr(err)
	minIdle, err := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE"))
	panicOnErr(err)

	Client = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URL"),
		Password:     os.Getenv("REDIS_PW"), // no password set
		DB:           0,                     // use default DB
		PoolSize:     poolSize,
		MinIdleConns: minIdle,
	})

	if _, err := Client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("package init ok: imredis")
}

func panicOnErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
