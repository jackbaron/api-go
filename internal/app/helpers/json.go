package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessPayload struct {
	Data    interface{}
	Message string
	Status  bool
}

func SendMessageSuccessFully(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	payload := SuccessPayload{
		Message: "success",
		Status:  true,
		Data:    data,
	}

	responseWithJson(w, 200, payload)
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
