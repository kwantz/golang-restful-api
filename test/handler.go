package test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kwantz/golang-restful-api/app"
	"github.com/kwantz/golang-restful-api/controller"
	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/repository"
	"github.com/kwantz/golang-restful-api/service"
)

func NewTestHandler(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, validate, categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)
	return app.NewRouter(categoryController)
}

func NewTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func TruncateTestDB(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func ToWebResponse(body []byte) map[string]interface{} {
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	return responseBody
}

func ToDataResponse(webResponse map[string]interface{}) map[string]interface{} {
	return webResponse["data"].(map[string]interface{})
}

func ToDataListResponse(webResponse map[string]interface{}) []interface{} {
	return webResponse["data"].([]interface{})
}
