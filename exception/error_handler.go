package exception

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kwantz/golang-restful-api/helper"
	"github.com/kwantz/golang-restful-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	} else if validationError(writer, request, err) {
		return
	} else {
		internalServerError(writer, request, err)
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		log.Printf("Validation Error: %+v\n", exception.Error())
		writer.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		})
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		log.Printf("Not Found Error: %+v\n", exception.Error)
		writer.WriteHeader(http.StatusNotFound)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		})
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	log.Printf("Internal Server Error: %+v\n", err)
	writer.WriteHeader(http.StatusInternalServerError)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	})
}
