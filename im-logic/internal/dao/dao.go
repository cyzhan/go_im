package dao

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	maria          *sql.DB
	UserDao        = &userDao{}
	VendorDao      = NewVendorDao()
	AbuseRecordDao = NewAbuseRecordDao()
	// elasticsearch
	ESMessageDao = NewESMessageDao()
)

func InitOldMaria() {
	var err error
	maria, err = sql.Open("mysql", os.Getenv("MARIA"))
	if err != nil {
		panic(err)
	}

	maria.SetConnMaxLifetime(3 * time.Second)
	maria.SetMaxOpenConns(10)
	maria.SetMaxIdleConns(10)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func offset(page int32, size int32) int32 {
	return (page - 1) * size
}

func GetTx(ctx context.Context) *sql.Tx {
	tx, err := maria.BeginTx(ctx, nil)
	panicOnErr(err)
	return tx
}

func exec(optTx *sql.Tx, script string, args ...any) sql.Result {
	var err error
	var res sql.Result
	if optTx != nil {
		res, err = optTx.Exec(script, args...)
	} else {
		res, err = maria.Exec(script, args...)
	}

	if err != nil && optTx != nil {
		optTx.Rollback()
		panicOnErr(err)
	} else if err != nil {
		panicOnErr(err)
	}

	return res
}
