package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var commands = map[string]func([]string, *sql.DB){
	"create":  create,
	"delete":  delete,
	"add":     add,
	"remove":  remove,
	"check":   check,
	"uncheck": uncheck,
	"show":    show,
}

func main() {
	db, err := sql.Open("sqlite3", "./todo.db")

	if err != nil {
		fmt.Println(fmt.Errorf("error creating connection to db %w", err))
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		PRAGMA foreign_keys = ON;
		CREATE TABLE IF NOT EXISTS lists (
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE
		);
		
		CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY,
		description TEXT UNIQUE,
		list_id INTEGER,
		checked INTEGER DEFAULT 0 NOT NULL,
		CONSTRAINT fk_lists
			FOREIGN KEY (list_id)
			REFERENCES lists(id)
			ON DELETE CASCADE
		);
	`)
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println(usage())
		os.Exit(2)
	}

	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Println(usage())
		os.Exit(2)
	} else {
		cmd(os.Args[2:], db)
	}
}
