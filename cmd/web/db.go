package main

import (
	"database/sql"
	"fmt"
)

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "i2MC1FMu7jlzQ"
	dbName = "snippetbox"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", getDBString(dbPort, dbHost, dbUser, dbPass, dbName))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getDBString(port int, host, user, password, dbname string) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
