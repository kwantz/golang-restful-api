package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kwantz/golang-restful-api/app"
	"github.com/kwantz/golang-restful-api/controller"
	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/repository"
	"github.com/kwantz/golang-restful-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, validate, categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Server run on port " + server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
