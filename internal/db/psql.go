package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log/slog"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "kanad4s"
	password = ""
	dbname   = "postgres"
)

type Response struct {
	Id     int    `db:"w_id"`
	Name   string `db:"w_name"`
	Salary int    `db:"w_salary"`
}

// var db *sqlx.DB

func Connect() (db *sqlx.DB) {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BD connected!\n")

	rows, err := db.Query("select * from professions")
	if err != nil {
		panic(err)
	}

	response := make([]Response, 0, 5)
	if err := sqlx.StructScan(rows, &response); err != nil {
		// log.Logger.Error("%s\n", err.Error())
	}
	fmt.Println(response)
	return
}

func Request(request string, db *sqlx.DB) {
	slog.Debug("Try request:" + request)
	rows, err := db.Query(request)
	if err != nil {
		panic(err)
	}
	response := make([]Response, 0, 5)
	if err := sqlx.StructScan(rows, &response); err != nil {
		// log.Logger.Error("%s\n", err.Error())
	}
	fmt.Println(response)
}

func Close(db *sqlx.DB) {
	db.Close()
}
