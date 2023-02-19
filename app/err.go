package app

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func ErrorJson(err error, code uint32) {
	if err != nil {
		LogWrite("error", fmt.Sprintf("json error %s %v\n", err.Error(), code))
		log.Printf("JSON Data Conversion error - %v\n", code)
	}
}

func ErrorTx(err error, tx *sqlx.Tx, debug ...interface{}) {
	if err != nil {
		LogWrite("error", fmt.Sprintf("tx error %s\n", err.Error()))
		err := tx.Rollback()
		if err != nil {
			LogWrite("error", fmt.Sprintf("tx rollback error %s\n", err.Error()))
		}
		return
	}
}

func ErrorPq(err error) {
	if err != nil {
		if err.(*pq.Error).Code == "23502" {
			LogWrite("error", err.Error())
		} else if err.(*pq.Error).Code == "23505" {
			LogWrite("error", err.Error())
		} else if err.(*pq.Error).Code == "42P01" {
			LogWrite("error", err.Error())
		} else {
			LogWrite("error", err.Error())
			panic(err.Error())
		}
	}
}

func ErrorFatal(err error) {
	if err != nil {
		LogWrite("error", fmt.Sprintf("fatal error %s\n", err.Error()))
		log.Fatal(err)
	}
}

func ErrorCheck(err error) {
	if err != nil {
		LogWrite("error", err.Error())
	}
}
