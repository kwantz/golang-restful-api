package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryControllerDelete(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Invalid name")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, GetCategoryDetailUrl(category.ID), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	handler := NewTestHandler(db)
	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	webReponse := ToWebResponse(body)

	assert.Equal(t, 204, response.StatusCode)
	assert.Equal(t, float64(204), webReponse["code"])
	assert.Equal(t, "No Content", webReponse["status"])
}

func BenchmarkCategoryControllerDelete(b *testing.B) {
	db := NewTestDB()
	TruncateTestDB(db)

	for i := 0; i < b.N; i++ {
		category := CreateNewCategory(db, "Invalid name")
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodDelete, GetCategoryDetailUrl(category.ID), nil)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "RAHASIA")

		handler := NewTestHandler(db)
		handler.ServeHTTP(recorder, request)
		recorder.Result()
	}
}
