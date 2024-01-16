package storage_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	data "github.com/teelowe/todo/data"
	s "github.com/teelowe/todo/storage"
)

// var db data.Database

// var db = s.SetupDB("sqlite3", "todo.db")

// func ClearTables() {
// 	_, err := db.Exec(`DELETE FROM lists`)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestMain(m *testing.M) {
	result := m.Run()
	//tear down db
	// db.Exec(`
	// PRAGMA writable_schema = 1;
	// delete from sqlite_master where type in ('table', 'index', 'trigger');
	// PRAGMA writable_schema = 0;sql
	// VACUUM;
	// PRAGMA INTEGRITY_CHECK;
	// `)
	os.Exit(result)
}
func TestInsertListNoError(t *testing.T) {
	err := s.InsertList("shopping", data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestInsertListWithError(t *testing.T) {
	firstErr := s.InsertList("packing", data.MockDB{E: true})
	assert.EqualError(t, firstErr, "error inserting list packing: MockDB Exec error")
}

func TestDeleteListNoError(t *testing.T) {
	deleteErr := s.DeleteList("shopping", data.MockDB{E: false})
	assert.Nil(t, deleteErr)
}

func TestDeleteListWithError(t *testing.T) {
	deleteErr := s.DeleteList("packing", data.MockDB{E: true})
	assert.EqualError(t, deleteErr, "error deleting list with name packing: MockDB Exec error")
}

func TestGetNameFromListNoError(t *testing.T) {
	err := s.GetNameFromList("foobar", data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestGetNameFromListWithError(t *testing.T) {
	err := s.GetNameFromList("foobar", data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "no list with name foobar exists")
}

func TestAddItemsNoError(t *testing.T) {
	list := "waawaaweewaa"
	err := s.AddItems([]string{"foo", "bar", "baz"}, "1", &list, data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestAddItemsWithError(t *testing.T) {
	list := "waawaaweewaa"
	err := s.AddItems([]string{"foo", "bar", "baz"}, "1", &list, data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error inserting new item into list waawaaweewaa: MockDB Exec error")
}

func TestDeleteItemsNoError(t *testing.T) {
	list := "woopeee"
	err := s.DeleteItems([]string{"foo"}, "1", &list, data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestDeleteItemsWithError(t *testing.T) {
	list := "woopeee"
	err := s.DeleteItems([]string{"foo"}, "1", &list, data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error removing item from list woopeee: MockDB Exec error")
}

func TestCheckItemsNoError(t *testing.T) {
	list := "woopeee"
	err := s.CheckItems([]string{"foo"}, &list, data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestCheckItemsWithError(t *testing.T) {
	list := "woopeee"
	err := s.CheckItems([]string{"foo"}, &list, data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error checking item foo in list woopeee: MockDB Exec error")
}

func TestUncheckItemsNoError(t *testing.T) {
	list := "woopeee"
	err := s.UncheckItems([]string{"foo"}, &list, data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestUncheckItemsWithError(t *testing.T) {
	list := "woopeee"
	err := s.UncheckItems([]string{"foo"}, &list, data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error unchecking item foo in list woopeee: MockDB Exec error")
}

func TestListExistsNoError(t *testing.T) {
	id, err := s.ListExists("foo", data.MockDB{E: false})
	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestListExistsWithError(t *testing.T) {
	id, err := s.ListExists("foo", data.MockDB{E: true})
	assert.Empty(t, id)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestItemExistsNoError(t *testing.T) {
	row, err := s.ItemExists("banana", data.MockDB{E: false})
	assert.Nil(t, err)
	assert.NotNil(t, row)
}

func TestItemExistsWithError(t *testing.T) {
	row, err := s.ItemExists("banana", data.MockDB{E: true})
	assert.Nil(t, row)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestGetSpecifiedListsNoError(t *testing.T) {
	rows, err := s.GetSpecifiedLists([]any{"foo", "bar"}, data.MockDB{E: false})
	assert.Nil(t, err)
	assert.NotNil(t, rows)
}

func TestGetSpecifiedListsWithError(t *testing.T) {
	rows, err := s.GetSpecifiedLists([]any{"foo", "bar"}, data.MockDB{E: true})
	assert.Nil(t, rows)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error selecting specific lists: MockDB Query error")
}

func TestGetAllListsNoError(t *testing.T) {
	rows, err := s.GetAllLists(data.MockDB{E: false})
	assert.Nil(t, err)
	assert.NotNil(t, rows)
}

func TestGetAllListsWithError(t *testing.T) {
	rows, err := s.GetAllLists(data.MockDB{E: true})
	assert.Nil(t, rows)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error selecting all lists: MockDB Query error")
}

func TestGetItemsByListIdNoError(t *testing.T) {
	rows, err := s.GetItemsByListId("1", data.MockDB{E: false})
	assert.Nil(t, err)
	assert.NotNil(t, rows)
}

func TestGetItemsByListIdWithError(t *testing.T) {
	rows, err := s.GetItemsByListId("1", data.MockDB{E: true})
	assert.NotNil(t, err)
	assert.Empty(t, rows)
	assert.EqualError(t, err, "error selecting list items: MockDB Query error")
}
