package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func InsertItems(items []string, listId string, list *string, db *sql.DB) {
	for _, i := range items {
		_, err := db.Exec(`
		INSERT INTO items (description, list_id) VALUES ($1, $2)`, i, listId)
		if err != nil {
			fmt.Println(fmt.Errorf("error inserting new item into list %s: %w", *list, err))
			os.Exit(1)
		}
		fmt.Printf("added item %s to list %s\n", i, *list)
	}
}

func CheckItems(items []string, list *string, db *sql.DB) {
	for i, v := range items {
		_, err := db.Exec(`
		UPDATE items SET CHECKED = $1 WHERE description = $2`, 1, items[i])
		if err != nil {
			fmt.Println(fmt.Errorf("error checking item %s in list %s: %w", v, *list, err))
			os.Exit(1)
		}
		fmt.Printf("checked item %s as done in list %s\n", v, *list)
	}
}

func CreateList(args []string, db *sql.DB) {
	for _, name := range args {
		_, err := db.Exec(`
		INSERT INTO lists (name) VALUES ($1)`, strings.ToLower(name))
		if err != nil {
			fmt.Println(fmt.Errorf("error inserting new list: %w", err))
			os.Exit(1)
		}
		fmt.Println("created list with name " + name)
	}
}

func DeleteList(args []string, db *sql.DB) {
	for _, name := range args {
		row := db.QueryRow("SELECT name FROM lists WHERE name = $1", strings.ToLower(name))
		var thisName string
		if err := row.Scan(&thisName); err == sql.ErrNoRows {
			fmt.Println(fmt.Errorf("no list with name %s exists", name))
			os.Exit(1)
		}
		_, err := db.Exec(`
		DELETE FROM lists WHERE name = ($1)`, strings.ToLower(name))
		if err != nil {
			fmt.Println(fmt.Errorf("error deleteing list with name %s: %w", name, err))
			os.Exit(1)
		}
		fmt.Println("deleted list with name " + name)
	}
}

func RemoveItemsFromList(items []string, listId string, list *string, db *sql.DB) {
	for _, i := range items {
		_, err := db.Exec(`
		DELETE FROM items WHERE description = $1 AND list_id = $2`, i, listId)
		if err != nil {
			fmt.Println(fmt.Errorf("error removing item from list %s: %w", *list, err))
			os.Exit(1)
		}
		fmt.Printf("removed item %s from list %s\n", i, *list)
	}
}

func GetListIdByName(name string, db *sql.DB) *sql.Row {
	return db.QueryRow("SELECT id FROM lists WHERE name = $1", name)
}

func GetItemDescription(description string, db *sql.DB) *sql.Row {
	return db.QueryRow("SELECT description FROM items WHERE items.description = $1", description)
}
