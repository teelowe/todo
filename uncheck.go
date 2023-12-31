package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

// uncheck a specified item(s) in a specified list
func uncheck(args []string, db *sql.DB) {
	uncheckCmd := flag.NewFlagSet("uncheck", flag.ExitOnError)
	list := uncheckCmd.String("l", "", "the name of the list containing the item to uncheck")
	item := uncheckCmd.String("i", "", "the name(s) of the item(s) to uncheck in the list")
	uncheckCmd.Parse(args)
	validateArgs(uncheckCmd, 2)
	exists, _ := listExists(*list, db)
	if exists {
		items := append([]string{*item}, uncheckCmd.Args()...)
		for _, v := range items {
			if itemExists(v, db) {
				uncheckItems(items, list, db)
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

func uncheckItems(items []string, list *string, db *sql.DB) {
	for i, v := range items {
		_, err := db.Exec(`
		UPDATE items SET CHECKED = $1 WHERE description = $2`, 0, items[i])
		if err != nil {
			fmt.Println(fmt.Errorf("error unchecking item %s in list %s: %w", v, *list, err))
			os.Exit(1)
		}
		fmt.Printf("unchecked item %s as not done in list %s\n", v, *list)
	}
}
