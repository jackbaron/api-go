package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nhatth/api-service/internal/app/controllers"
)

func authRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", controllers.Login)

	return r
}
