package main

import (
	"NotesServer/internal/cli"
	"NotesServer/internal/db"
	_ "NotesServer/internal/log"
	"fmt"
	"log/slog"
)

func init() {
	// log.Setup()
}

func main() {
	dbInfo := db.DbInfo{
		Host:    "localhost",
		Port:    5432,
		Name:    "postgres",
		Sslmode: "disable",
	}
	db1 := db.Connect(dbInfo)
	db.Query("select * from professions", db1)
	// db.Query("", db1)
	db.Close(db1)
	cli.ShowNavigation()
	val, err := cli.GetNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

	slog.Info("Finishing...")
}
