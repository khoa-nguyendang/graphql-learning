package main

import (
	"database/sql"
	"fmt"
	"server/shared"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func GetDB(c *Config) *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.Dbname)
	db, err := sql.Open("postgres", dbinfo)
	shared.CheckErr(err)
	return db
}

func RunMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	shared.CheckErr(err)
	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	shared.CheckErr(err)
	m.Up()
}
