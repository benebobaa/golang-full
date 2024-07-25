package json

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
	err := json.NewDecoder(request.Body).Decode(result)
	return err
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	json.NewEncoder(writer).Encode(response)
}
