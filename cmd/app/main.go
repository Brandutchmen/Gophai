package main

import (
	"app/internal/config"
	"app/internal/database"
	"app/internal/router"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	database.Connect()
	router.SetupRoutes(r)
	fmt.Println("Server is running on port http://localhost:3000")
	fmt.Println("GraphQL playground is running on http://localhost:3000/api/graphql/playground")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
