package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nhatth/api-service/internal/app/composer"
)

func AuthRouter(authService composer.AuthService) http.Handler {
	r := chi.NewRouter()

	r.Post("/login", authService.LoginHdl)

	r.Post("/regitser", authService.RegisterHdl)

	r.Post("/token/refresh", authService.RefreshTokenHl)

	return r
}
