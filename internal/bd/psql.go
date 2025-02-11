package bd

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

// type res struct {
// 	id int
//     name string
//     salary int
// }

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
	rows, ok := db.Query("select * from professions")
	if ok != nil {
		panic(err)
	}
	// fmt.Printf("response: \n%v\n", response)
	
	// fmt.Println(response.)
	ids := make([]int, 0)
	names := make([]string, 0)
	salaries := make([]int, 0)
    for rows.Next() {
        var id int
        var name string
        var salary int
		// var resp res
        if err := rows.Scan(&id, &name, &salary); err != nil {
            log.Fatal(err)
        }
		// fmt.Println(resp)
        ids = append(ids, id)
        names = append(names, name)
        salaries = append(salaries, salary)
    }
	fmt.Println(ids)
	fmt.Println(names)
	fmt.Println(salaries)
}