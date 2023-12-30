package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

// show will display a specified list if provided, or all lists if no -l flag was provided
func show(args []string, db *sql.DB) {
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	list := showCmd.String("l", "", "the name of the list to show (optional)")
	showCmd.Parse(args)
	allProvidedLists := clean(append([]string{*list}, showCmd.Args()...))
	listFlagProvided := showCmd.Lookup("l").Value.String()
	var lists map[string]string
	if listFlagProvided == "" {
		lists = getLists(nil, db)
	} else {
		lists = getLists(allProvidedLists, db)
	}
	var tdl TodoList
	var tdls TodoLists
	for k, v := range lists {
		tdl = todoList(k, v, db)
		tdls = append(tdls, tdl)
	}
	renderOutput(tdls)
}

// render the todo lists in table format
func renderOutput(data TodoLists) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"List Name", "Todo Item", "Status"})
	for _, v := range data {
		t.AppendRow([]interface{}{v.Name})
		for _, q := range v.Items {
			t.AppendRow([]interface{}{"", q.Description, renderStatus(q.Checked)})
		}
	}
	t.SetStyle(table.StyleLight)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignCenter},
	})
	t.Render()
}

// render a green check mark for "done", and empty box for "not done"
func renderStatus(checked int) string {
	if checked == 1 {
		return string("\u2705")
	}
	return string("\u25A1")
}

// return map[id]name for specified lists or all lists
func getLists(listNames []string, db *sql.DB) map[string]string {
	var queryListIds *sql.Rows
	var err error
	if listNames != nil {
		queryListIds, err = db.Query("SELECT id, name FROM lists WHERE name in ($1)", strings.Join(listNames[:], ","))
	} else {
		queryListIds, err = db.Query("SELECT id, name FROM lists")
	}
	if err != nil {
		fmt.Println(fmt.Errorf("error selecting lists: %w", err))
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

// return a TodoList based on a list id and a list name
func todoList(id string, name string, db *sql.DB) TodoList {
	todoItems := make([]TodoItem, 0)
	todoList := TodoList{}
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
	if td.Err() != nil {
		panic(td.Err())
	}
	return todoList
}
