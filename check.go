package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

// check an item(s) in a specified list (i.e. mark it as "done")
func check(args []string, db *sql.DB) {
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	list := checkCmd.String("l", "", "the name of the list containing the item to check")
	item := checkCmd.String("i", "", "the name(s) of the item(s) to check in the list")
	checkCmd.Parse(args)
	validateArgs(checkCmd, 2)

	_, err := listExists(*list, db)
	if err != nil {
		fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist: %w", *list, err))
		os.Exit(1)
	}
	items := append([]string{*item}, checkCmd.Args()...)
	checkItems(items, list, db)
}

func checkItems(items []string, list *string, db *sql.DB) {
	for i, v := range items {
		_, err := db.Exec(`
		UPDATE items SET CHECKED = $1 WHERE description = $2`, 1, items[i])
		if err != nil {
			fmt.Println(fmt.Errorf("error checking item %s in list %s: %w", v, *list, err))
			os.Exit(1)
		}
		fmt.Printf("checked item %s as done in list %s\n", v, *list)
	}
}
