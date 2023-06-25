package idgenerator

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"imws/internal/util/imredis"

	"github.com/google/uuid"
)

const (
	ID_MIN = 100000
)

func Get(ctx context.Context, chatID int64) (int64, bool, error) {
	increase := rand.Int63n(3) + 1
	msgID, err := imredis.Client.IncrBy(ctx, imredis.KeyOfMsgID(chatID), increase).Result()
	if err != nil {
		return 0, false, err
	}

	if msgID >= ID_MIN {
		return msgID, false, nil
	}

	uuidStr, isRequired := requireLock(ctx, imredis.KeyOfMsgIDLock(chatID), 5)
	if isRequired {
		firstMsgID := ID_MIN + rand.Int63n(3) + 1
		imredis.Client.Set(ctx, imredis.KeyOfMsgID(chatID), strconv.FormatInt(firstMsgID, 10), 90*24*time.Hour).Result()
		imredis.Client.Del(ctx, uuidStr)
		return firstMsgID, true, nil
	}

	for i := 0; i < 5; i++ {
		time.Sleep(20 * time.Millisecond)
		msgID, err := imredis.Client.IncrBy(ctx, imredis.KeyOfMsgID(chatID), increase).Result()
		if err != nil {
			return 0, false, err
		}

		if msgID >= ID_MIN {
			return msgID, false, nil
		}
	}

	return 0, false, errors.New("something goes wrong requiring new message id")
}

func requireLock(ctx context.Context, key string, expiredSec uint32) (string, bool) {
	uuidStr := uuid.New().String()
	isSuccess, err := imredis.Client.SetNX(ctx, key, uuidStr, 5*time.Second).Result()
	if isSuccess {
		return uuidStr, true
	}
	if err != nil {
		log.Printf(err.Error())
	}
	return "", false
}
