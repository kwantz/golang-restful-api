package helper

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ReadInt64FromParams(params httprouter.Params, name string) int64 {
	param := params.ByName(name)
	id, err := strconv.ParseInt(param, 10, 32)
	PanicIfError(err)
	return int64(id)
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
