package http_utils

import (
	"encoding/json"
	"github.com/theguptaji/bookstore_utils-go/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, StatusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJson(w, err.Status(), err)
}
