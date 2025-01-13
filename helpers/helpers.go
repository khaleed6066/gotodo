package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JsonResponse sends a JSON response for both success and error scenarios.
func JsonResponse(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: success,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func HandleError(w http.ResponseWriter, err error, message string, statusCode int) {
	if err != nil {
		JsonResponse(w, statusCode, false, message, nil)
	}
}
