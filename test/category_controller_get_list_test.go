package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kwantz/golang-restful-api/model/entity"
	"github.com/stretchr/testify/assert"
)

func TestCategoryControllerGetList(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	categories := []entity.Category{
		CreateNewCategory(db, "First Category"),
		CreateNewCategory(db, "Second Category"),
		CreateNewCategory(db, "Third Category"),
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, GetCategoryUrl(), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	handler := NewTestHandler(db)
	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	webReponse := ToWebResponse(body)
	categoriesResponse := ToDataListResponse(webReponse)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, float64(200), webReponse["code"])
	assert.Equal(t, "OK", webReponse["status"])

	for idx := range categoriesResponse {
		categoryResponse := categoriesResponse[idx].(map[string]interface{})
		assert.Equal(t, categories[idx].Name, categoryResponse["name"])
		assert.Equal(t, float64(categories[idx].ID), categoryResponse["id"])
	}
}

func BenchmarkCategoryControllerGetList(b *testing.B) {
	db := NewTestDB()
	TruncateTestDB(db)
	CreateNewCategory(db, "First Category")
	CreateNewCategory(db, "Second Category")
	CreateNewCategory(db, "Third Category")
	CreateNewCategory(db, "Forth Category")
	CreateNewCategory(db, "Fifth Category")

	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, GetCategoryUrl(), nil)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "RAHASIA")

		handler := NewTestHandler(db)
		handler.ServeHTTP(recorder, request)
		recorder.Result()
	}
}
