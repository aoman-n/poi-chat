package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/resolvers"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic recovered: %v", r)
		}
	}()

	fmt.Printf("conf: %+v \n\n", config.Env)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	r.Use(cors.Default().Handler)
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
