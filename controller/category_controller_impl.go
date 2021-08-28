package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/model/web"
	"github.com/kwantz/golang-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	category := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   category,
	})
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryUpdateRequest.ID = helper.ReadInt64FromParams(params, "categoryID")
	category := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	})
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := helper.ReadInt64FromParams(params, "categoryID")
	controller.CategoryService.Delete(request.Context(), categoryID)
	writer.WriteHeader(http.StatusNoContent)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusNoContent,
		Status: "No Content",
	})
}

func (controller CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := helper.ReadInt64FromParams(params, "categoryID")
	category := controller.CategoryService.FindByID(request.Context(), categoryID)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	})
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll(request.Context())
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categories,
	})
}
