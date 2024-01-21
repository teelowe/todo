package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
	"github.com/teelowe/todo/util"
)

// check an item(s) in a specified list (i.e. mark it as "done")
func check(args []string, db data.Database) {
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	list := checkCmd.String("l", "", "the name of the list containing the item to check")
	item := checkCmd.String("i", "", "the name(s) of the item(s) to check in the list")
	checkCmd.Parse(args)
	if checkCmd.NFlag() != 2 {
		fmt.Println("please provide the -l and -i flags")
		os.Exit(1)
	}
	_, err := storage.ListExists(*list, db)
	if err != nil {
		fmt.Printf("the provided list %s doesn't exist\n", *list)
		os.Exit(1)
	}
	items := util.Clean(append([]string{*item}, checkCmd.Args()...))
	for _, v := range items {
		_, err := storage.ItemExists(v, db)
		if err != nil {
			fmt.Printf("the specified item '%s' doesn't exist in list %s\n", v, *list)
			os.Exit(1)
		}
	}
	storage.CheckItems(items, list, db)
}
