package main

import (
	"NotesServer/internal/cli"
	"NotesServer/internal/db"
	"NotesServer/internal/log"
	"fmt"
	"log/slog"
)

func init() {
	log.Setup()
	// log.Info("Main init")
	// slog.Info("abs", slog.String("arg1", "val1"), slog.String("arg2", "val2"))
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
	db.Close(db1)
	cli.ShowNavigation()
	val, err := cli.GetNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	slog.Info("Finishing...")
}
