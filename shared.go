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

// return true if the provided list name matches an existing list, else false
func listExists(list_name string, db *sql.DB) (bool, string) {
	row := db.QueryRow("SELECT id FROM lists WHERE name = $1", list_name)
	var thisId string
	if err := row.Scan(&thisId); err == sql.ErrNoRows {
		return false, ""
	}
	return true, thisId
}

// return true if the provided lists_id
func itemExists(item_name string, db *sql.DB) bool {
	row := db.QueryRow("SELECT description FROM items WHERE items.description = $1", item_name)
	var thisId string
	if err := row.Scan(&thisId); err == sql.ErrNoRows {
		return false
	}
	return true
}

// return s with items lowercased and without leading/trailing commas
func clean(s []string) []string {
	var cleaned []string
	for _, v := range s {
		cleaned = append(cleaned, strings.ToLower(strings.Trim(v, ",")))
	}
	return cleaned
}

// print usage for the app
func usage() string {
	s := "Usage: todo [command] [flags]\nAvailable commands:\n"
	for k := range commands {
		s += " - " + k + "\n"
	}
	return s
}
