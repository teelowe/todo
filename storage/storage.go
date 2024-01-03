package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	DB *sql.DB
}

func NewStore(driver, dsn string) (*Storage, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening connection to db %w", err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &Storage{db}, nil
}

func (s *Storage) Close() {
	s.DB.Close()
}
