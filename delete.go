package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
)

func delete(args []string, db *sql.DB) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.String("n", "", "the name of the list to delete")
	err := deleteCmd.Parse(args)
	if err != nil {
		fmt.Println(err)
	}

	for _, name := range args[1:] {
		row := db.QueryRow("SELECT name FROM lists WHERE name = $1", strings.ToLower(name))
		var thisName string
		if err := row.Scan(&thisName); err == sql.ErrNoRows {
			fmt.Println(fmt.Errorf("no list with name %s exists", name))
			os.Exit(1)
		}
		_, err = db.Exec(`
		DELETE FROM lists WHERE name = ($1)`, strings.ToLower(name))
		if err != nil {
			fmt.Println(fmt.Errorf("error deleteing list with name %s: %w", name, err))
			os.Exit(1)
		}
		fmt.Println("deleted list with name " + name)
	}
}
