package main

import (
	"flag"
	"fmt"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

// delete a list and all of its associated tasks
func delete(args []string, db data.Database) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.String("l", "", "the name of the list to delete")
	err := deleteCmd.Parse(args)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range args[1:] {
		if err := storage.GetNameFromList(v, db); err != nil {
			fmt.Println(err)
		}
		if err = storage.DeleteList(v, db); err != nil {
			fmt.Println(err)
		}
	}
}
