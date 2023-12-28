package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
)

func create(args []string, db *sql.DB) {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.String("n", "", "the name of the list to create")
	createCmd.Parse(args)
	validateArgs(createCmd, 1)

	for _, name := range args[1:] {
		_, err := db.Exec(`
		INSERT INTO lists (name) VALUES ($1)`, strings.ToLower(name))
		if err != nil {
			fmt.Println(fmt.Errorf("error inserting new list: %w", err))
			os.Exit(1)
		}
		fmt.Println("created list with name " + name)
	}
}
