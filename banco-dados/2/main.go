package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1500.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1500.00},
	// 	{Name: "Teclado", Price: 250.00},
	// 	{Name: "Monitor", Price: 800.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Teclado")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// select all with limit and offset
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%te%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	var p Product
	db.First(&p, 1)
	p.Name = "New Notebook"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2)

	db.Delete(&p2)
}
