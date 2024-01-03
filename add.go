package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/storage"
)

// add a specified item (todo) to a specified list
func add(args []string, db *sql.DB) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	list := addCmd.String("l", "", "the name of the list to add items to")
	item := addCmd.String("i", "", "the name(s) of the item(s) to add to the list")
	addCmd.Parse(args)
	validateArgs(addCmd, 2)
	exists, id := listExists(*list, db)
	if exists {
		items := clean(append([]string{*item}, addCmd.Args()...))
		storage.InsertItems(items, id, list, db)
		return
	}
	fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist", *list))
	os.Exit(1)
}
