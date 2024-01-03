package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/teelowe/todo/storage"
)

// throw error if number of provided args is less than numArgs
func validateArgs(flagset *flag.FlagSet, numArgs int) {
	if flagset.NFlag() != numArgs {
		fmt.Println(fmt.Errorf("missing required number of flags"))
		os.Exit(2)
	}
}

// return true and listid if the provided list name matches an existing list, else false and ""
func listExists(list_name string, db *sql.DB) (bool, string) {
	row := storage.GetListIdByName(list_name, db)
	var thisId string
	if err := row.Scan(&thisId); err == sql.ErrNoRows {
		return false, ""
	}
	return true, thisId
}

// return true if the provided lists_id
func itemExists(item_name string, db *sql.DB) bool {
	row := storage.GetItemDescription(item_name, db)
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
