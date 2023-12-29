package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed todo.db
var content embed.FS

func getEmbeddedFileContent() ([]byte, error) {
	return content.ReadFile("todo.db")
}

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
	dbFilePath := "todo.db"

	// Check if the database file already exists
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		// If the file doesn't exist, create it and write the embedded content
		data, err := getEmbeddedFileContent()
		if err != nil {
			fmt.Println("Error reading embedded file:", err)
			return
		}

		if err := os.WriteFile(dbFilePath, data, 0644); err != nil {
			fmt.Println("Error creating database file:", err)
			return
		}

		fmt.Println("Database file created successfully.")
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		fmt.Println(fmt.Errorf("error opening connection to db %w", err))
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
