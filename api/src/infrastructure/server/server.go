package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/delivery/graphql"
	"github.com/laster18/poi/api/src/delivery/rest"
	"github.com/laster18/poi/api/src/infrastructure/db"
	"github.com/laster18/poi/api/src/infrastructure/redis"
	customMiddleware "github.com/laster18/poi/api/src/middleware"
	"github.com/laster18/poi/api/src/repository"
	"github.com/rs/cors"
)

func Init() {
	db := db.NewDb()
	redisClient := redis.New(config.Conf.Redis)
	roomRepo := repository.NewRoomRepo(db)

	router := chi.NewRouter()
	resolver := graphql.NewResolver(db, redisClient)
	conf := generated.Config{Resolvers: resolver}

	rooms, err := roomRepo.ListAll(context.Background())
	if err != nil {
		log.Println(err)
	}

	for _, r := range rooms {
		go resolver.SetupRoom(r.ID)
		fmt.Printf("[done setup] room: %+v \n", r)
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.GetHead)
	router.Use(middleware.Recoverer)
	router.Use(customMiddleware.AuthMiddleware())

	srv := handler.New(generated.NewExecutableSchema(conf))

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	// rest
	rest.NewRoutes(router)

	// graphql
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// serve image files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Conf.Port)
	log.Fatal(http.ListenAndServe(":"+config.Conf.Port, router))
}
