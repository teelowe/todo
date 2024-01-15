package main

import (
	"flag"
	"fmt"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

// create a new list with a given name
func create(args []string, db data.Database) {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.String("l", "", "the name of the list to create")
	createCmd.Parse(args)
	for _, v := range args[1:] {
		err := storage.InsertList(v, db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
