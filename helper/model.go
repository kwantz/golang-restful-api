package helper

import (
	"github.com/kwantz/golang-restful-api/model/entity"
	"github.com/kwantz/golang-restful-api/model/web"
)

func ToCategoryResponse(category entity.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []entity.Category) []web.CategoryResponse {
	responses := []web.CategoryResponse{}
	for _, category := range categories {
		responses = append(responses, ToCategoryResponse(category))
	}
	return responses
}
