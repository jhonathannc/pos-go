package main

import (
	"context"
	"database/sql"

	"github.com/jhonathannc/pos-go/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "Backend",
	// 	Description: sql.NullString{
	// 		String: "Backend description",
	// 		Valid:  true,
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:   "8760245f-331a-4c25-b1a2-c8db8659df32",
	// 	Name: "Backend Updated",
	// 	Description: sql.NullString{
	// 		String: "Backend description updated",
	// 		Valid:  true,
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	err = queries.DeleteCategory(ctx, "8760245f-331a-4c25-b1a2-c8db8659df32")
	if err != nil {
		panic(err)
	}
}
