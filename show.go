package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/teelowe/todo/data"
	"github.com/teelowe/todo/storage"
	"github.com/teelowe/todo/util"
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
func show(args []string, db data.Database) {
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	list := showCmd.String("l", "", "the name of the list to show (optional)")
	showCmd.Parse(args)
	allProvidedLists := util.Clean(append([]string{*list}, showCmd.Args()...))
	listFlagProvided := showCmd.Lookup("l").Value.String()
	var lists map[string]string
	var err error
	if listFlagProvided == "" {
		lists, err = getLists(nil, db)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		lists, err = getLists(allProvidedLists, db)
		if err != nil {
			fmt.Println(err)
		}
	}
	var tdl TodoList
	var tdls TodoLists
	for k, v := range lists {
		tdl, err = todoList(k, v, db)
		if err != nil {
			fmt.Println(err)
		}
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
func getLists(listNames []string, db data.Database) (map[string]string, error) {
	var queryListIds data.RowsInterface
	var err error
	if listNames != nil {
		args := make([]any, len(listNames))
		for i, list := range listNames {
			args[i] = list
		}
		queryListIds, err = storage.GetSpecifiedLists(args, db)
		if err != nil {
			return nil, err
		}
	} else {
		queryListIds, err = storage.GetAllLists(db)
		if err != nil {
			return nil, err
		}
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
	return lists, nil
}

// return a TodoList based on a list id and a list name
func todoList(id string, name string, db data.Database) (TodoList, error) {
	todoItems := make([]TodoItem, 0)
	todoList := TodoList{}
	td, err := storage.GetItemsByListId(id, db)
	if err != nil {
		return TodoList{}, err
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
	return todoList, nil
}
