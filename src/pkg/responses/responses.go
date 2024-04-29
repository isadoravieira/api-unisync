package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// DomainJSON returns a JSON response to request
func DomainJSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(datas); err != nil {
		log.Fatal(err)
	}
}

// DomainError returns an error in JSON format
func DomainError(w http.ResponseWriter, statusCode int, erro error) {
	DomainJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: erro.Error(),
	})
}