package middleware

import (
	"net/http"

	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) http.Handler {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") == "RAHASIA" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		})
	}
}
