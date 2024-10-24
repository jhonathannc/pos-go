package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	useCase := NewUseCase(db)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}
