package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nhatth/api-service/internal/app/composer"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/api/v1", v1(db))

	return r
}

func v1(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	authService := composer.ComposeAuthAPIService(db)

	r.Mount("/auth", authRouter(authService))

	return r
}
