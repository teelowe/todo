package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/teelowe/todo/data"
)

func SetupDB(driver, dsn string) *data.SQLiteDatabase {
	// if db file doesn't exist, create one
	if _, err := os.Stat(dsn); errors.Is(err, fs.ErrNotExist) {
		if err := os.WriteFile(dsn, []byte{}, 0644); err != nil {
			log.Fatal("Error creating database file:", err)
		}
		fmt.Println("database file created successfully")
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatalf("sql.Open(driver, dsn) failed with '%s'\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db.Ping() failed with '%s'\n", err)
	}
	_, err = db.Exec(`
		PRAGMA foreign_keys = ON;
		CREATE TABLE IF NOT EXISTS lists (
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE
		);
		
		CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY,
		description TEXT UNIQUE,
		list_id INTEGER,
		checked INTEGER DEFAULT 0 NOT NULL,
		CONSTRAINT fk_lists
			FOREIGN KEY (list_id)
			REFERENCES lists(id)
			ON DELETE CASCADE
		);`,
	)
	if err != nil {
		log.Fatalf("db.Exec() failed with '%s'\n", err)
	}
	return &data.SQLiteDatabase{Connection: db}
}

func InsertList(arg string, db data.Database) error {
	const query = `INSERT INTO lists (name) VALUES ($1)`
	_, err := db.Exec(query, strings.ToLower(arg))
	if err != nil {
		return fmt.Errorf("error inserting list %v: %v", arg, err)
	} else {
		fmt.Printf("created list with name %v\n", arg)
	}
	return nil
}

func DeleteList(name string, db data.Database) error {
	const deleteQuery = `DELETE FROM lists WHERE name = $1`
	_, err := db.Exec(deleteQuery, strings.ToLower(name))
	if err != nil {
		return fmt.Errorf("error deleting list with name %s: %w", name, err)
	}
	fmt.Println("deleted list with name " + name)
	return nil
}

func GetNameFromList(name string, db data.Database) error {
	const selectQuery = `SELECT name FROM lists WHERE name = $1`
	row := db.QueryRow(selectQuery, strings.ToLower(name))
	var thisName string
	if err := row.Scan(&thisName); err == sql.ErrNoRows {
		return fmt.Errorf("no list with name %s exists", name)
	}
	return nil
}

func AddItems(items []string, listId string, list *string, db data.Database) error {
	query := `INSERT INTO items (description, list_id) VALUES ($1, $2)`
	for _, v := range items {
		_, err := db.Exec(query, v, listId)
		if err != nil {
			return fmt.Errorf("error inserting new item into list %s: %w", *list, err)
		}
		fmt.Printf("added item %s to list %s\n", v, *list)
	}
	return nil
}

func DeleteItems(items []string, listId string, list *string, db data.Database) error {
	query := `DELETE FROM items WHERE description = $1 AND list_id = $2`
	for _, i := range items {
		_, err := db.Exec(query, i, listId)
		if err != nil {
			return fmt.Errorf("error removing item from list %s: %w", *list, err)
		}
		fmt.Printf("removed item %s from list %s\n", i, *list)
	}
	return nil
}

func CheckItems(items []string, list *string, db data.Database) error {
	query := `UPDATE items SET checked = $1 WHERE description = $2`
	for i, v := range items {
		_, err := db.Exec(query, 1, items[i])
		if err != nil {
			return fmt.Errorf("error checking item %s in list %s: %w", v, *list, err)
		}
		fmt.Printf("checked item %s as done in list %s\n", v, *list)
	}
	return nil
}

func UncheckItems(items []string, list *string, db data.Database) error {
	query := `UPDATE items SET checked = $1 WHERE description = $2`
	for i, v := range items {
		_, err := db.Exec(query, 0, items[i])
		if err != nil {
			return fmt.Errorf("error unchecking item %s in list %s: %w", v, *list, err)
		}
		fmt.Printf("unchecked item %s as not done in list %s\n", v, *list)
	}
	return nil
}

func ListExists(list_name string, db data.Database) (string, error) {
	row := db.QueryRow(`SELECT id FROM lists WHERE name = $1`, list_name)
	var thisId string
	if err := row.Scan(&thisId); err == sql.ErrNoRows {
		return "", err
	}
	return thisId, nil
}

func ItemExists(item_name string, db data.Database) (data.RowInterface, error) {
	row := db.QueryRow(`SELECT description FROM items WHERE items.description = $1`, item_name)
	var thisItem string
	if err := row.Scan(&thisItem); err == sql.ErrNoRows {
		return nil, err
	}
	return row, nil
}

func GetSpecifiedLists(args []any, db data.Database) (data.RowsInterface, error) {
	query := `SELECT id, name FROM lists WHERE name in (?` + strings.Repeat(",?", len(args)-1) + `)`
	queryListIds, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error selecting specific lists: %w", err)
	}
	return queryListIds, nil
}

func GetAllLists(db data.Database) (data.RowsInterface, error) {
	query := `SELECT id, name FROM lists`
	var args []any
	queryListIds, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error selecting all lists: %w", err)
	}
	return queryListIds, nil
}

func GetItemsByListId(id string, db data.Database) (data.RowsInterface, error) {
	query := `SELECT description, checked FROM items WHERE list_id = $1`
	todos, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error selecting list items: %w", err)
	}
	return todos, nil
}
