package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func StartPostgres() *sql.DB {
	host := "localhost"
	port := "8881"
	user := "nbcamp-user"
	pass := "nbcamp-pass"
	dbname := "nbcamp-be"

	dbc, err := getPostgres(host, port, user, pass, dbname)
	if err != nil {
		panic(err)
	}

	return dbc

}

func getPostgres(host, port, user, pass, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	return createConnection(dsn)
}

func createConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
