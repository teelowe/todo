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
func TestCreateListNoError(t *testing.T) {
	err := s.InsertList("shopping", data.MockDB{E: false})
	assert.Nil(t, err)
}

func TestCreateListWithError(t *testing.T) {
	firstErr := s.InsertList("packing", data.MockDB{E: true})
	assert.EqualError(t, firstErr, "error inserting list shopping: test error")
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
	err := s.GetNameFromList("foobar", data.MockDB{E: false})
	assert.Nil(t, err)
}
