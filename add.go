package main

import (
	"flag"
	"fmt"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
	"github.com/teelowe/todo/util"
)

// add a specified item (todo) to a specified list
func add(args []string, db data.Database) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	list := addCmd.String("l", "", "the name of the list to add items to")
	item := addCmd.String("i", "", "the name(s) of the item(s) to add to the list")
	err := addCmd.Parse(args)
	if err != nil {
		fmt.Println(err)
	}
	id, err := storage.ListExists(*list, db)
	if err != nil {
		fmt.Println(err)
	}
	items := util.Clean(append([]string{*item}, addCmd.Args()...))
	err = storage.AddItems(items, id, list, db)
	if err != nil {
		fmt.Println(err)
	}
}
