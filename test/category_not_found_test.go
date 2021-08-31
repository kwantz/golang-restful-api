package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CategoryNotFoundRequest struct {
	Method string
	Name   string
	Body   io.Reader
}

func TestCategoryControllerGetNotFound(t *testing.T) {
	db := NewTestDB()
	TruncateTestDB(db)

	requests := []CategoryNotFoundRequest{
		{Name: "NotFoundCategoryGetRequest", Method: http.MethodGet, Body: nil},
		{Name: "NotFoundCategoryUpdateRequest", Method: http.MethodPut, Body: strings.NewReader(`{"name" : "Movie"}`)},
		{Name: "NotFoundCategoryDeleteRequest", Method: http.MethodDelete, Body: nil},
	}
	for _, request := range requests {
		t.Run(request.Name, func(t *testing.T) {
			newRecorder := httptest.NewRecorder()
			newRequest := httptest.NewRequest(request.Method, GetCategoryDetailUrl(123), request.Body)
			newRequest.Header.Add("Content-Type", "application/json")
			newRequest.Header.Add("X-API-Key", "RAHASIA")

			handler := NewTestHandler(db)
			handler.ServeHTTP(newRecorder, newRequest)

			response := newRecorder.Result()
			body, _ := io.ReadAll(response.Body)
			webReponse := ToWebResponse(body)

			assert.Equal(t, 404, response.StatusCode)
			assert.Equal(t, float64(404), webReponse["code"])
			assert.Equal(t, "Not Found", webReponse["status"])
			assert.Equal(t, "category not found", webReponse["data"])
		})
	}
}
