package main

import (
	"app/database"
	"app/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	database.Connect()
	router.SetupRoutes(r)
	http.ListenAndServe(":3000", r)
}
