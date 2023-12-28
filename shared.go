package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
)

// throw error if number of provided args is less than numArgs
func validateArgs(flagset *flag.FlagSet, numArgs int) {
	if flagset.NFlag() != numArgs {
		fmt.Println(fmt.Errorf("missing required number of flags"))
		os.Exit(2)
	}
}

// if the list exists, return its id and nil error, else empty list and error
func listExists(name string, db *sql.DB) (string, error) {
	row := db.QueryRow("SELECT id FROM lists WHERE name = $1", strings.ToLower(name))
	var thisId string
	if err := row.Scan(&thisId); err == sql.ErrNoRows {
		return "", err
	} else {
		return thisId, nil
	}
}

func usage() string {
	s := "Usage: todo [command] [flags]\nAvailable commands:\n"
	for k := range commands {
		s += " - " + k + "\n"
	}
	return s
}
