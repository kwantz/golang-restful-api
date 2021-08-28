package repository

import (
	"context"
	"database/sql"

	"github.com/kwantz/golang-restful-api/model/entity"
)

type CategoryRepository interface {
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category)
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Create(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int64) (entity.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
