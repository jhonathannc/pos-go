package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jhonathannc/pos-go/uow/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS `courses`;")
	dbt.Exec("DROP TABLE IF EXISTS `categories`;")

	dbt.Exec("CREATE TABLE categories (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, description varchar(255));")
	dbt.Exec("CREATE TABLE courses (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id int NOT NULL);")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
