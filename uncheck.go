package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

// uncheck a specified item(s) in a specified list
func uncheck(args []string, db data.Database) {
	uncheckCmd := flag.NewFlagSet("uncheck", flag.ExitOnError)
	list := uncheckCmd.String("l", "", "the name of the list containing the item to uncheck")
	item := uncheckCmd.String("i", "", "the name(s) of the item(s) to uncheck in the list")
	uncheckCmd.Parse(args)
	if uncheckCmd.NFlag() != 2 {
		fmt.Println("please provide the -l and -i flags")
		os.Exit(1)
	}
	_, err := storage.ListExists(*list, db)
	if err != nil {
		fmt.Printf("the provided list %s doesn't exist\n", *list)
	}
	items := append([]string{*item}, uncheckCmd.Args()...)
	for _, v := range items {
		_, err := storage.ItemExists(v, db)
		if err != nil {
			fmt.Printf("the provided item %s doesn't exist in the provided list %s\n", v, *list)
			os.Exit(1)
		}
	}
	err = storage.UncheckItems(items, list, db)
	if err != nil {
		fmt.Println(err)
	}
}
