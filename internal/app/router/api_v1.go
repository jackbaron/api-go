package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/api/v1", v1())

	return r
}

func v1() http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", authRouter())

	return r
}
