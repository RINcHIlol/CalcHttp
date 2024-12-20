package calculation

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func ResponseError(w http.ResponseWriter, err string, statusCode int) {
	errorData := Error{Error: err}
	w.WriteHeader(statusCode)
	errorJson, _ := json.Marshal(errorData)
	w.Write(errorJson)
}
