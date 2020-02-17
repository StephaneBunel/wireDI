package main

import (
	"database/sql"
	"fmt"
)

// ProvideDB returns a new DB connection according to the URI
func ProvideDB(uri string) (*sql.DB, func(), error) {

	var cleanup = func() {
		fmt.Println("cleanup DB")
	}

	return nil, cleanup, nil
}
