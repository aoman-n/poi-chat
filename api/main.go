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
	"github.com/laster18/poi/api/src/delivery/graphql"
	"github.com/laster18/poi/api/src/delivery/rest"
	"github.com/laster18/poi/api/src/infrastructure"
	"github.com/laster18/poi/api/src/repository"
	"github.com/rs/cors"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic recovered: %v", r)
		}
	}()

	fmt.Printf("conf: %+v \n\n", config.Conf)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.Conf.Port
	}

	db := infrastructure.NewDb()
	_ = repository.NewRoomRepository(db)

	r := chi.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	r.Use(cors.Default().Handler)

	// rest
	rest.NewRoutes(r)

	// graphql
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
