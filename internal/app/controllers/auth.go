package controllers

import (
	"net/http"

	"github.com/nhatth/api-service/internal/app/helpers"
)

func Login(w http.ResponseWriter, r *http.Request) {
	helpers.SendMessageSuccessFully(w, r, 200, struct{}{})
}
