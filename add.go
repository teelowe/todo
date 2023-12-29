package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
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
		insertItems(items, id, list, db)
		return
	}

	fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist", *list))
	os.Exit(1)
}

func insertItems(items []string, listId string, list *string, db *sql.DB) {
	for _, i := range items {
		_, err := db.Exec(`
		INSERT INTO items (description, list_id) VALUES ($1, $2)`, i, listId)
		if err != nil {
			fmt.Println(fmt.Errorf("error inserting new item into list %s: %w", *list, err))
			os.Exit(1)
		}
		fmt.Printf("added item %s to list %s\n", i, *list)
	}
}
