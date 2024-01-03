package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/teelowe/todo/storage"
)

// delete a list and all of its associated tasks
func delete(args []string, db *sql.DB) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.String("l", "", "the name of the list to delete")
	err := deleteCmd.Parse(args)
	if err != nil {
		fmt.Println(err)
	}
	storage.DeleteList(args, db)
}
