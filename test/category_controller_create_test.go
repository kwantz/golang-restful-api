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

func TestCategoryControllerCreate(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, GetCategoryUrl(), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	handler := NewTestHandler(db)
	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	webReponse := ToWebResponse(body)
	categoryResponse := ToDataResponse(webReponse)

	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, float64(201), webReponse["code"])
	assert.Equal(t, "Created", webReponse["status"])
	assert.Equal(t, "Gadget", categoryResponse["name"])
}

func TestCategoryControllerCreateValidation(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, GetCategoryUrl(), requestBody)
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

func BenchmarkCategoryControllerCreate(b *testing.B) {
	db := NewTestDB()
	TruncateTestDB(db)

	for i := 0; i < b.N; i++ {
		requestBody := strings.NewReader(fmt.Sprintf(`{"name" : "Category %d"}`, i))
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, GetCategoryUrl(), requestBody)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-API-Key", "RAHASIA")

		handler := NewTestHandler(db)
		handler.ServeHTTP(recorder, request)
		recorder.Result()
	}
}
