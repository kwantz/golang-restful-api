package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kwantz/golang-restful-api/controller"
	"github.com/kwantz/golang-restful-api/exception"
	"github.com/kwantz/golang-restful-api/middleware"
)

func NewRouter(categoryController controller.CategoryController) http.Handler {
	router := httprouter.New()
	setupCategoryRouter(router, categoryController)

	router.PanicHandler = exception.ErrorHandler
	return middleware.NewAuthMiddleware(router)
}

func setupCategoryRouter(router *httprouter.Router, categoryController controller.CategoryController) {
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)

	router.GET("/api/categories/:categoryID", categoryController.FindByID)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)
}
