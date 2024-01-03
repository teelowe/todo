package main

import (
	"database/sql"
	"flag"

	"github.com/teelowe/todo/storage"
)

// create a new list with a given name
func create(args []string, db *sql.DB) {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.String("l", "", "the name of the list to create")
	createCmd.Parse(args)
	validateArgs(createCmd, 1)
	storage.CreateList(args[1:], db)
}
