package bd

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	host = "localhost"
	port     = 5432
    user     = "postgres"
    password = "your-password"
    dbname   = "calhounio_demo"
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BD connected!")
}