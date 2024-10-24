package main

import (
	"database/sql"

	"github.com/jhonathannc/pos-go/di/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := product.NewProductRepository(db)
	useCase := product.NewProductUseCase(repository)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}
