package main

import (
	"flag"
	"fmt"

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
	id, err := storage.ListExists(*list, db)
	if err != nil {
		fmt.Println(err)
	}
	items := util.Clean(append([]string{*item}, removeCmd.Args()...))
	err = storage.DeleteItems(items, id, list, db)
	if err != nil {
		fmt.Println(err)
	}
}
