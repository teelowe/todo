package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

func uncheck(args []string, db *sql.DB) {
	uncheckCmd := flag.NewFlagSet("uncheck", flag.ExitOnError)
	list := uncheckCmd.String("l", "", "the name of the list containing the item to uncheck")
	item := uncheckCmd.String("i", "", "the name(s) of the item(s) to uncheck in the list")
	uncheckCmd.Parse(args)
	validateArgs(uncheckCmd, 2)

	_, err := listExists(*list, db)
	if err != nil {
		fmt.Println(fmt.Errorf("the provided list '%s' doesn't exist: %w", *list, err))
		os.Exit(1)
	}
	items := append([]string{*item}, uncheckCmd.Args()...)
	uncheckItems(items, list, db)
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
