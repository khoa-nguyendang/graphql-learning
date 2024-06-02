package main

import (
	"database/sql"
	"fmt"
	"server/shared"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "Pass1234!"
	DB_NAME     = "postgres"
)

func GetDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	shared.CheckErr(err)
	return db
}
