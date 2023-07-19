package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nhatth/api-service/internal/app/composer"
	"github.com/nhatth/api-service/internal/app/database"
	"github.com/nhatth/api-service/internal/app/helpers"
	authRoutes "github.com/nhatth/api-service/internal/app/services/auth/routes"
	"gorm.io/gorm"
)

func main() {

	//? Load config file
	config, err := helpers.LoadConfig("./../")

	if err != nil {

		log.Fatalln("Cannot load config file")

		return
	}

	//? Connection DB
	db := database.ConnectDatabase(config)

	r := setUpRoutes(db)

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		log.Println("Running server faild")
	}

	log.Println("Running server port 8000")
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
