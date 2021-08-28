package test

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/kwantz/golang-restful-api/model/entity"
	"github.com/kwantz/golang-restful-api/repository"
)

func CreateNewCategory(db *sql.DB, name string) entity.Category {
	tx, _ := db.Begin()
	category := entity.Category{Name: name}
	categoryRepository := repository.NewCategoryRepository()
	category = categoryRepository.Create(context.Background(), tx, category)
	tx.Commit()
	return category
}

func GetCategoryUrl() string {
	return "http://localhost:2000/api/categories"
}

func GetCategoryDetailUrl(id int64) string {
	return GetCategoryUrl() + "/" + strconv.FormatInt(id, 10)
}
