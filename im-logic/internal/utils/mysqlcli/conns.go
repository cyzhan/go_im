package mysqlcli

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mariaConfig struct {
	Url      string
	Account  string
	Password string
	Database string
	LogMod   bool
}

type IMySQLClient interface {
	Session() *gorm.DB
	Close() error
}

type DBClient struct {
	Client *gorm.DB
}

func initWithConfig(config mariaConfig) *DBClient {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8", config.Account, config.Password, config.Url, config.Database,
	)

	var err error
	db, err := gorm.Open(mysql.Open(connect))
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", connect, err))
	}

	if config.LogMod {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("Database [%s] Connect success", config.Database))
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	return &DBClient{Client: db}
}

func (c *DBClient) Session() *gorm.DB {
	db := c.Client.Session(&gorm.Session{})
	return db
}

func (c *DBClient) Close() error {
	client, err := c.Client.DB()
	if err != nil {
		return err
	}

	if err = client.Close(); err != nil {
		return err
	}

	return nil
}
