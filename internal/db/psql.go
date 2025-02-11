package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

//  http://84.237.50.81:8080/apex

const (
	host 	 = "localhost"
	port     = 5432
    user     = "kanad4s"
    password = ""
    dbname   = "postgres"
)

type Response struct {
	id int
    name string
    salary int
}

func Connect() {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
        host, port, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BD connected!\n")
	rows, err := db.Query("select * from professions")
	if err != nil {
		panic(err)
	}
	
	response := make([]Response, 0)
    for rows.Next() {
        var id int
        var name string
        var salary int
        if err := rows.Scan(&id, &name, &salary); err != nil {
            log.Fatal(err)
        }
        response = append(response, Response{
			id: id,
			name: name,
			salary: salary,
		})
    }
	fmt.Println(response)
}