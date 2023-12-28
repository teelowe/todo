package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type TodoLists []TodoList

type TodoList struct {
	Name  string
	Items []TodoItem
}

type TodoItem struct {
	Description string
	Checked     int
}

func show(args []string, db *sql.DB) {
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	list := showCmd.String("l", "", "the name of the list to show (optional)")

	showCmd.Parse(args)
	// validateArgs(showCmd, 1)
	// listFlagProvided := showCmd.Lookup("l").Value.String()
	// if listFlagProvided == "" {
	// 	// show all lists
	// } else {
	// 	// show specified list(s)
	// }
	allProvidedLists := append([]string{*list}, showCmd.Args()...)
	fmt.Println(allProvidedLists)

	l := allLists(db)
	var tdl TodoList
	var tdls TodoLists
	for k, v := range l {
		tdl = todoList(k, v, db)
		tdls = append(tdls, tdl)
	}
	// fmt.Printf("%+v\n", tdls)
	b, _ := json.MarshalIndent(tdls, "", "   ")
	fmt.Println(string(b))
}

func allLists(db *sql.DB) map[string]string {
	queryListIds, err := db.Query("SELECT id, name FROM lists")
	if err != nil {
		fmt.Println(fmt.Errorf("error selecting all lists: %w", err))
		os.Exit(1)
	}
	defer queryListIds.Close()
	lists := make(map[string]string, 0)
	for queryListIds.Next() {
		var id, name string
		if err := queryListIds.Scan(&id, &name); err != nil {
			panic(err)
		}
		lists[id] = name
	}
	if queryListIds.Err() != nil {
		panic(queryListIds.Err())
	}
	return lists
}

func todoList(id string, name string, db *sql.DB) TodoList {
	todoItems := make([]TodoItem, 0)
	todoList := TodoList{}
	// todoLists := TodoLists{}
	td, err := db.Query("SELECT description, checked FROM items WHERE list_id = $1", id)
	if err != nil {
		fmt.Println(fmt.Errorf("error selecting list items: %w", err))
		os.Exit(1)
	}
	for td.Next() {
		var todo TodoItem
		if err := td.Scan(&todo.Description, &todo.Checked); err != nil {
			panic(err)
		}
		todoItems = append(todoItems, todo)
	}
	todoList.Name = name
	todoList.Items = todoItems
	// todoLists = append(todoLists, todoList)
	if td.Err() != nil {
		panic(td.Err())
	}
	return todoList
}
