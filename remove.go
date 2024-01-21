package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
	"github.com/teelowe/todo/util"
)

// remove specified item(s) from a specified list
func remove(args []string, db data.Database) {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	list := removeCmd.String("l", "", "the name of the list to remove items from")
	item := removeCmd.String("i", "", "the name(s) of the item(s) to remove from the list")
	removeCmd.Parse(args)
	if removeCmd.NFlag() != 2 {
		fmt.Println("please provide the -l and -i flags")
		os.Exit(1)
	}
	id, err := storage.ListExists(*list, db)
	if err != nil {
		fmt.Printf("the provided list %s doesn't exist\n", *list)
		os.Exit(1)
	}
	items := util.Clean(append([]string{*item}, removeCmd.Args()...))
	err = storage.DeleteItems(items, id, list, db)
	if err != nil {
		fmt.Println(err)
	}
}
