package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

// delete a list and all of its associated tasks
func delete(args []string, db data.Database) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.String("l", "", "the name of the list to delete")
	deleteCmd.Parse(args)
	if deleteCmd.NFlag() != 1 {
		fmt.Println("please provide the -l flag and specify a list name")
		os.Exit(1)
	}
	for _, v := range args[1:] {
		if err := storage.GetNameFromList(v, db); err != nil {
			fmt.Println(err)
		}
		if err := storage.DeleteList(v, db); err != nil {
			fmt.Println(err)
		}
	}
}
