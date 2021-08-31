package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UnauthorizedRequest struct {
	Method string
	Name   string
	URL    string
}

func TestUnauthorizationWithoutAPIKey(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	requests := []UnauthorizedRequest{
		{Name: "Unauthorized Get List Request", Method: http.MethodGet, URL: GetCategoryUrl()},
		{Name: "Unauthorized Create Request", Method: http.MethodPost, URL: GetCategoryUrl()},
		{Name: "Unauthorized Get Request", Method: http.MethodGet, URL: GetCategoryDetailUrl(123)},
		{Name: "Unauthorized Update Request", Method: http.MethodPut, URL: GetCategoryDetailUrl(123)},
		{Name: "Unauthorized Delete Request", Method: http.MethodDelete, URL: GetCategoryDetailUrl(123)},
	}
	for _, request := range requests {
		t.Run(request.Name, func(t *testing.T) {
			newRecorder := httptest.NewRecorder()
			newRequest := httptest.NewRequest(request.Method, request.URL, nil)
			handler := NewTestHandler(db)
			handler.ServeHTTP(newRecorder, newRequest)

			response := newRecorder.Result()
			body, _ := io.ReadAll(response.Body)
			webReponse := ToWebResponse(body)

			assert.Equal(t, 401, response.StatusCode)
			assert.Equal(t, float64(401), webReponse["code"])
			assert.Equal(t, "Unauthorized", webReponse["status"])
		})
	}
}

func TestUnauthorizationWrongAPIKey(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	requests := []UnauthorizedRequest{
		{Name: "Unauthorized Get List Request", Method: http.MethodGet, URL: GetCategoryUrl()},
		{Name: "Unauthorized Create Request", Method: http.MethodPost, URL: GetCategoryUrl()},
		{Name: "Unauthorized Get Request", Method: http.MethodGet, URL: GetCategoryDetailUrl(123)},
		{Name: "Unauthorized Update Request", Method: http.MethodPut, URL: GetCategoryDetailUrl(123)},
		{Name: "Unauthorized Delete Request", Method: http.MethodDelete, URL: GetCategoryDetailUrl(123)},
	}
	for _, request := range requests {
		t.Run(request.Name, func(t *testing.T) {
			newRecorder := httptest.NewRecorder()
			newRequest := httptest.NewRequest(request.Method, request.URL, nil)
			newRequest.Header.Add("X-API-Key", "AISAHAR")

			handler := NewTestHandler(db)
			handler.ServeHTTP(newRecorder, newRequest)

			response := newRecorder.Result()
			body, _ := io.ReadAll(response.Body)
			webReponse := ToWebResponse(body)

			assert.Equal(t, 401, response.StatusCode)
			assert.Equal(t, float64(401), webReponse["code"])
			assert.Equal(t, "Unauthorized", webReponse["status"])
		})
	}
}
