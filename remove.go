package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/storage"
)

// remove specified item(s) from a specified list
func remove(args []string, db *sql.DB) {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	list := removeCmd.String("l", "", "the name of the list to remove items from")
	item := removeCmd.String("i", "", "the name(s) of the item(s) to remove from the list")
	removeCmd.Parse(args)
	validateArgs(removeCmd, 2)

	exists, id := listExists(*list, db)
	if exists {
		items := clean(append([]string{*item}, removeCmd.Args()...))
		storage.RemoveItemsFromList(items, id, list, db)
		return
	}
	fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist", *list))
	os.Exit(1)
}
