package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/model/entity"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.ID)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "INSERT INTO categories (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.ID = id
	return category
}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int64) (entity.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		category := scanCategory(ctx, rows)
		return category, nil
	} else {
		return entity.Category{}, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []entity.Category{}
	for rows.Next() {
		category := scanCategory(ctx, rows)
		categories = append(categories, category)
	}
	return categories
}

func scanCategory(ctx context.Context, rows *sql.Rows) entity.Category {
	category := entity.Category{}
	err := rows.Scan(&category.ID, &category.Name)
	helper.PanicIfError(err)
	return category
}
