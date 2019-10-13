package httperror

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	HttpCode int    `json:"HttpCode"`
	Message  string `json:"Message" bson:"Message"`
}

func New(httpCode int, message string, rw http.ResponseWriter) *Error {
	errorModel := &Error{httpCode, message}
	if rw != nil {
		errorMessage, _ := json.Marshal(&errorModel)
		rw.WriteHeader(httpCode)
		rw.Write(errorMessage)
	}
	return errorModel
}
