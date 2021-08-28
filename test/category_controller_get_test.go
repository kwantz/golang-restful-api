package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryControllerGet(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Food")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, GetCategoryDetailUrl(category.ID), nil)
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
	assert.Equal(t, category.Name, categoryResponse["name"])
	assert.Equal(t, float64(category.ID), categoryResponse["id"])
}

func BenchmarkCategoryControllerGet(b *testing.B) {
	db := NewTestDB()
	TruncateTestDB(db)

	category := CreateNewCategory(db, "Food")
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, GetCategoryDetailUrl(category.ID), nil)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "RAHASIA")

		handler := NewTestHandler(db)
		handler.ServeHTTP(recorder, request)
		recorder.Result()
	}
}
