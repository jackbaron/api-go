package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nhatth/api-service/internal/app/composer"
	"github.com/nhatth/api-service/internal/app/database"
	"github.com/nhatth/api-service/internal/app/helpers"
	authRoutes "github.com/nhatth/api-service/internal/app/services/auth/routes"
	"github.com/nhatth/api-service/pkg/logger"
	"gorm.io/gorm"
)

func main() {

	//? Load config file
	absPath, err := filepath.Abs("./../")
	if err != nil {
		panic(err)
	}

	config, err := helpers.LoadConfig(absPath)

	logger := logger.GlobalLogger().GetLogger("service")

	if err != nil {

		logger.Fatalf("Cannot load config file %s", err.Error())

		return
	}

	//? Connection DB
	db := database.ConnectDatabase(config)

	r := setUpRoutes(db)

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Running server port 8000")
}

func setUpRoutes(db *gorm.DB) *chi.Mux {

	authService := composer.ComposeAuthAPIService(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/api/v1", v1(authService))

	return r
}

func v1(authService composer.AuthService) http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", authRoutes.AuthRouter(authService))

	return r
}
