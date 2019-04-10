package result

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteErrorResponse(responseWriter http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(responseWriter, msgbytes, http.StatusBadRequest)
}

func WriteJsonResponse(responseWriter http.ResponseWriter, bytes []byte, code int) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(code)
	_, err := responseWriter.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
