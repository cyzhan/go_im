package imredis

import (
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func init() {
	poolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	checkErr(err)
	minIdle, err := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE"))
	checkErr(err)

	Client = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URL"),
		Password:     os.Getenv("REDIS_PW"), // no password set
		DB:           0,                     // use default DB
		PoolSize:     poolSize,
		MinIdleConns: minIdle,
	})

	log.Println("package init ok: imredis")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func KeyOfMsgID(chatID int64) string {
	return fmt.Sprintf("msgID-generator:%d", chatID)
}

func KeyOfMsgIDLock(chatID int64) string {
	return fmt.Sprintf("msgID-lock:%d", chatID)
}

func KeyOfChat(chatID int64) string {
	return fmt.Sprintf("chat:%d", chatID)
}

func KeyOfGroup(groupId int64) string {
	return fmt.Sprintf("group:%d", groupId)
}
