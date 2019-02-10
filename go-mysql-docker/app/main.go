package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {

	http.HandleFunc("/products", func(rw http.ResponseWriter, r *http.Request) {
		db := connectDb()
		rows, err := db.Query("SELECT * FROM products")
		if err != nil {
			panic(err)
		}

		var products []Product

		for rows.Next() {
			var product Product
			if err := rows.Scan(&product.ID,
				&product.Name,
				&product.Description); err != nil {
				// handle error
			}

			products = append(products, product)
		}

		response, _ := json.Marshal(products)

		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(200)
		rw.Write(response)
	})

	http.ListenAndServe(":8080", nil)
}

func connectDb() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "dlmbg", "123456", "db", "3306", "db_dummy"))
	if err != nil {
		panic("Could not connect to the db")
	} else {
		return db
	}
}
