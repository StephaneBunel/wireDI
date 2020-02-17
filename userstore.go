package main

import (
	"database/sql"
	"fmt"
)

// UserStoreConfig is the configuration passed to ProvideUserStore
type UserStoreConfig struct{}

// UserStore is the handle to access User data
type UserStore struct {
	pool *sql.DB
	conf UserStoreConfig
}

// ProvideUserStoreDefaultConfig returns a default configuration to use with ProvideUserStore()
func ProvideUserStoreDefaultConfig() UserStoreConfig {
	return UserStoreConfig{}
}

// ProvideUserStore returns a new UserStore that uses db and conf as dependencies
func ProvideUserStore(db *sql.DB, conf UserStoreConfig) (*UserStore, func(), error) {
	var us = new(UserStore)
	us.conf = conf
	us.pool = db

	var cleanup = func() {
		fmt.Println("cleanup user store")
	}

	return us, cleanup, nil
}
