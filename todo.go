package main

import (
	"fmt"
	"os"

	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
)

var commands = map[string]func(args []string, db data.Database){
	"create":  create,
	"delete":  delete,
	"add":     add,
	"remove":  remove,
	"check":   check,
	"uncheck": uncheck,
	"show":    show,
}

func main() {
	db := storage.SetupDB("sqlite3", "todo.db")
	defer db.Close()
	if len(os.Args) < 2 {
		fmt.Println(usage())
		os.Exit(2)
	}

	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Println(usage())
		os.Exit(2)
	} else {
		cmd(os.Args[2:], db)
	}
}

// print usage for the app
func usage() string {
	s := "Usage: todo [command] [flags]\nAvailable commands:\n"
	for k := range commands {
		s += " - " + k + "\n"
	}
	return s
}
