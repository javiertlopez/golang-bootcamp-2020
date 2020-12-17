package controller

import (
	"encoding/json"
	"net/http"
)

// Response struct
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// JSONResponse writter
func JSONResponse(w http.ResponseWriter, code int, response interface{}) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Convert our interface to JSON
	output, err := json.Marshal(response)

	if err != nil {
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(output)
}
