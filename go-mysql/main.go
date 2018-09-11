package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare statement for reading data
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("%d : %s \n", id, name)
	}
}
