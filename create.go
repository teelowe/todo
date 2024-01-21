package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

// create a new list with a given name
func create(args []string, db data.Database) {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.String("l", "", "the name of the list to create")
	createCmd.Parse(args)
	if createCmd.NFlag() != 1 {
		fmt.Println("please provide the -l flag and specify a list name")
		os.Exit(1)
	}
	for _, v := range args[1:] {
		err := storage.InsertList(v, db)
		if err != nil {
			fmt.Println(err)
		}
	}
}
