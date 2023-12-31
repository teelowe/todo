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
	exists, _ := listExists(*list, db)
	if exists {
		items := clean(append([]string{*item}, checkCmd.Args()...))
		for _, v := range items {
			if itemExists(v, db) {
				checkItems(items, list, db)
			} else {
				fmt.Println(fmt.Errorf("the specified item '%s' doesn't exist in list %s", v, *list))
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist", *list))
		os.Exit(1)
	}
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
