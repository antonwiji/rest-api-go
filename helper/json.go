package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(write http.ResponseWriter, respon interface{}) {

	write.Header().Add("Content-Type", "application-json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(respon)
	PanicIfError(err)
}
