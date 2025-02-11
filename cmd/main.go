package main

import (
	// "fmt"
    "NotesServer/internal/db"
)

func main() {
	db1 := db.Connect()
    db.Request("select * from professions", db1)
    db.Close(db1)
    // fmt.Println(bd)
}