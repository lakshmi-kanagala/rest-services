package main

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "Missing@2022"
	DB_NAME     = "postgres"
	host        = "localhost"
	port        = 5432
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Println("DB info ::" + dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	v, ok := db.Exec("SELECT * FROM employee.Employee")
	fmt.Println("pinging to DB", db.Ping())
	fmt.Println(ok)
	fmt.Println(v)
	checkErr(err)

	return db
}
