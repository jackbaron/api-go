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

type SuccessOnlyMessage struct {
	Message string
	Status  bool
}

func SendMessageSuccessWithOutPayLoad(w http.ResponseWriter, r *http.Request, code int) {
	payload := SuccessOnlyMessage{
		Message: "success",
		Status:  true,
	}

	responseWithJson(w, code, payload)
}

func SendMessageSuccessFully(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	payload := SuccessPayload{
		Message: "success",
		Status:  true,
		Data:    data,
	}

	responseWithJson(w, code, payload)
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
