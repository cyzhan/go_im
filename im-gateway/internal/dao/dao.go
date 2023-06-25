package dao

import (
	"context"
	"database/sql"
	"time"

	"log"

	"imgateway/internal/model/entity"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	maria  *sql.DB
	Vendor = &vendorDao{
		tokenVdIDMap: make(map[string]*entity.VendorEntity),
	}
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitMaria() {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		err, ok := r.(error)
		if ok {
			log.Fatalln(err.Error())
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

	list := Vendor.GetVendors()
	for _, v := range list {
		Vendor.tokenVdIDMap[v.Token] = v
		log.Printf("vdID:%v, token:%s", v.ID, v.Token)
	}

	log.Println("package init ok: dao")
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
