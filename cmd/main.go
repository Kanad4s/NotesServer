package main

import (
	"fmt"
    "NotesServer/internal/db"
    "NotesServer/internal/cli"
    "NotesServer/internal/log"
)

func main() {
	db1 := db.Connect()
    db.Request("select * from professions", db1)
    db.Close(db1)
    cli.ShowNavigation()
    val, err := cli.GetNumber()
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
    log.Logger.Info("Finishing...")
}