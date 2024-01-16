package data

import (
	"database/sql"
	"errors"
)

// data is the abstraction for the database allowing mocking of the calls to the database/sql package.
// This was an interesting exercise in mocking, but I wouldn't recommend it.  It's a lot of cognitive overhead
// for the purposes of easing testing of the database calls, but for such a simple application, this is silly.
// Probably would have made more sense to just use a test db.  Then again, it was educational, and surprisingly...
// it WORKS!!!!  Go is pretty cool if you can get your head around it :)

type Database interface {
	Exec(query string, args ...any) (*sql.Result, error)
	QueryRow(query string, args ...any) RowInterface
	Query(query string, args ...any) (RowsInterface, error)
	Close() error
}

type SQLiteDatabase struct {
	Connection *sql.DB
}

func (s *SQLiteDatabase) Exec(query string, args ...any) (*sql.Result, error) {
	result, err := s.Connection.Exec(query, args...)
	return &result, err
}

func (s *SQLiteDatabase) QueryRow(query string, args ...any) RowInterface {
	row := s.Connection.QueryRow(query, args...)
	return row
}

func (s *SQLiteDatabase) Query(query string, args ...any) (RowsInterface, error) {
	rows, err := s.Connection.Query(query, args...)
	return CustomRows{rows}, err
}

func (s *SQLiteDatabase) Close() error {
	s.Connection.Close()
	return nil
}

type MockDB struct {
	E bool
}

func (m MockDB) Exec(query string, args ...any) (*sql.Result, error) {
	var res *sql.Result
	if m.E {
		return res, errors.New("MockDB Exec error")
	}
	return res, nil
}

func (m MockDB) QueryRow(query string, args ...any) RowInterface {
	if m.E {
		return &MockSqlRow{E: true}
	}
	return &MockSqlRow{}
}

func (m MockDB) Query(query string, args ...any) (RowsInterface, error) {
	if m.E {
		return &MockSqlRows{E: true}, errors.New("MockDB Query error")
	}
	return &MockSqlRows{}, nil
}

func (m MockDB) Close() error {
	return nil
}

type RowInterface interface {
	Scan(destination ...any) error
	Err() error
}

type MockSqlRow struct {
	E bool
}

func (msr MockSqlRow) Err() error {
	if msr.E {
		return errors.New("mock sql row error")
	}
	return nil
}

func (msr MockSqlRow) Scan(destination ...any) error {
	if msr.E {
		return sql.ErrNoRows
	}
	return nil
}

type RowsInterface interface {
	Next() bool
	Err() error
	Close() error
	Scan(dest ...any) error
}

type MockSqlRows struct {
	E bool
}

func (msrs MockSqlRows) Next() bool {
	return false
}

func (msrs MockSqlRows) Err() error {
	return nil
}

func (msrs MockSqlRows) Close() error {
	return nil
}

func (msrs MockSqlRows) Scan(dest ...any) error {
	return nil
}

// CustomRows is a custom type that embeds *sql.Rows and implements RowsInterface
type CustomRows struct {
	*sql.Rows
}

func (cr CustomRows) GetUnderlyingRows() *sql.Rows {
	return cr.Rows
}
