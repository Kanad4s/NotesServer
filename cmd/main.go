package main

import (
	"NotesServer/internal/cli"
	"NotesServer/internal/db"
	"fmt"
	// "NotesServer/internal/log"
	"log/slog"
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
	slog.Info("Finishing...")
}
