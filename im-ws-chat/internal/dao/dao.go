package dao

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"imws/internal/model/errcode"
)

var (
	maria        *sql.DB
	Vendor       *vendorDao
	User         *userDao
	UserTimeline *userTimelineDao
)

func init() {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		err, ok := r.(error)
		if ok {
			log.Fatalf(err.Error())
		} else {
			log.Println("recover with unknown type")
			os.Exit(1)
		}
	}()

	var err error
	maria, err = sql.Open("mysql", os.Getenv("MARIA"))
	if err != nil {
		panic(err)
	}
	maria.SetConnMaxLifetime(time.Minute * 3)
	maria.SetMaxOpenConns(10)
	maria.SetMaxIdleConns(10)
	log.Println("package init ok: dao")
}

func panicIfError(err error) {
	if err != nil {
		panic(errcode.InternalError(err))
	}
}

func GetTx(ctx context.Context) *sql.Tx {
	tx, err := maria.BeginTx(ctx, nil)
	panicIfError(err)
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
		panic(errcode.InternalError(err))
	} else if err != nil {
		panic(errcode.InternalError(err))
	} else {
		return res
	}
}
