package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RESPONSE(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		RESPONSE(w, statusCode, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	RESPONSE(w, http.StatusBadRequest, nil)
}