package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryControllerUpdate(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Film")
	requestBody := strings.NewReader(`{"name" : "Movie"}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, GetCategoryDetailUrl(category.ID), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	handler := NewTestHandler(db)
	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	webReponse := ToWebResponse(body)
	categoryResponse := ToDataResponse(webReponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, float64(200), webReponse["code"])
	assert.Equal(t, "OK", webReponse["status"])
	assert.Equal(t, "Movie", categoryResponse["name"])
	assert.Equal(t, float64(category.ID), categoryResponse["id"])
}

func TestCategoryControllerUpdateValidation(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Film")
	requestBody := strings.NewReader(`{"name" : ""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, GetCategoryDetailUrl(category.ID), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	handler := NewTestHandler(db)
	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	webReponse := ToWebResponse(body)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, float64(400), webReponse["code"])
	assert.Equal(t, "Bad Request", webReponse["status"])
}

func BenchmarkCategoryControllerUpdate(b *testing.B) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Film")
	for i := 0; i < b.N; i++ {
		requestBody := strings.NewReader(fmt.Sprintf(`{"name" : "Movie %d"}`, i))
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut, GetCategoryDetailUrl(category.ID), requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "RAHASIA")

		handler := NewTestHandler(db)
		handler.ServeHTTP(recorder, request)
		recorder.Result()
	}
}
