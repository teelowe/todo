package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

// remove a specified item(s) from a specified list
func remove(args []string, db *sql.DB) {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	list := removeCmd.String("l", "", "the name of the list to remove items from")
	item := removeCmd.String("i", "", "the name(s) of the item(s) to remove from the list")
	removeCmd.Parse(args)
	validateArgs(removeCmd, 2)

	listId, err := listExists(*list, db)
	if err != nil {
		fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist: %w", *list, err))
		os.Exit(1)
	}
	items := append([]string{*item}, removeCmd.Args()...)
	removeItems(items, listId, list, db)
}

func removeItems(items []string, listId string, list *string, db *sql.DB) {
	for _, i := range items {
		_, err := db.Exec(`
		DELETE FROM items WHERE description = $1 AND list_id = $2`, i, listId)
		if err != nil {
			fmt.Println(fmt.Errorf("error removing item from list %s: %w", *list, err))
			os.Exit(1)
		}
		fmt.Printf("removed item %s from list %s\n", i, *list)
	}
}
