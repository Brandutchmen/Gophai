package router

import (
	"app/internal/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/graphql", func(r chi.Router) {
			// TODO: playground should be disabled in production
			r.Handle("/playground", playground.Handler("GraphQL playground", "/api/graphql"))
			r.Handle("/", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))
		})
	})
}
