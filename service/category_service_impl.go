package service

import (
	"context"
	"database/sql"

	"github.com/kwantz/golang-restful-api/exception"
	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/model/entity"
	"github.com/kwantz/golang-restful-api/model/web"
	"github.com/kwantz/golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	DB                 *sql.DB
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(db *sql.DB, categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		DB:                 db,
		CategoryRepository: categoryRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := entity.Category{Name: request.Name}
	category = service.CategoryRepository.Create(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := service.findCategoryByIdOrPanicNotFound(ctx, tx, request.ID)
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID int64) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := service.findCategoryByIdOrPanicNotFound(ctx, tx, categoryID)
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID int64) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	category := service.findCategoryByIdOrPanicNotFound(ctx, tx, categoryID)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoriesResponse(categories)
}

func (service *CategoryServiceImpl) findCategoryByIdOrPanicNotFound(ctx context.Context, tx *sql.Tx, categoryID int64) entity.Category {
	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	exception.PanicNotFoundIfError(err)
	return category
}
