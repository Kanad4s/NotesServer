package main

import (
	"NotesServer/internal/cli"
	"NotesServer/internal/db"
	"NotesServer/internal/log"
	"fmt"
	"log/slog"
)

func main() {
	db1 := db.Connect()
	db.Request("select * from professions", db1)
	defer db.Close(db1)
	cli.ShowNavigation()
	val, err := cli.GetNumber()
	if err != nil {
		panic(err)
	}
	// log.Logger.Info("Finishing...")
	fmt.Println(val)
	log.TestAll()
	slog.Debug("Finishing...")
}
