package mysqlcli

import (
	"os"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once   sync.Once
	client *DBClient
)

func NewDBClient() IMySQLClient {
	logMod, _ := strconv.ParseBool(os.Getenv("MARIA_LOG_MOD"))
	config := mariaConfig{
		Url:      os.Getenv("MARIA_URL"),
		Account:  os.Getenv("MARIA_ACCOUNT"),
		Password: os.Getenv("MARIA_PW"),
		Database: os.Getenv("MARIA_DATABASE"),
		LogMod:   logMod,
	}

	once.Do(func() {
		client = initWithConfig(config)
	})

	return client
}

func GetClient() *DBClient {
	return client
}
