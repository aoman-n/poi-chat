package server

import (
	"context"
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
	"github.com/laster18/poi/api/src/infra/db"
	"github.com/laster18/poi/api/src/infra/redis"
	"github.com/laster18/poi/api/src/presentation/graphql/directive"
	customMiddleware "github.com/laster18/poi/api/src/presentation/graphql/middleware"
	"github.com/laster18/poi/api/src/presentation/graphql/resolver"
	"github.com/laster18/poi/api/src/presentation/rest"
	"github.com/laster18/poi/api/src/registry"
	"github.com/rs/cors"
)

func Start() {
	ctx := context.Background()
	db := db.NewDb()
	redisClient := redis.New()

	repo := registry.NewRepository(db, redisClient)
	svc := registry.NewService(repo)
	subscriber := registry.NewSubscriber(redisClient, repo.NewUser())

	userSubscriber := subscriber.NewUser()
	roomUserSubscriber := subscriber.NewRoomUser()
	go userSubscriber.Start(ctx)
	go roomUserSubscriber.Start(ctx)

	router := chi.NewRouter()
	resolver := resolver.New(repo, svc, roomUserSubscriber, userSubscriber)
	directive := directive.New(repo)
	conf := generated.Config{Resolvers: resolver, Directives: *directive}

	// set middlewares
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8080",
		},
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
	router.Use(middleware.RealIP)
	router.Use(middleware.GetHead)
	router.Use(middleware.Recoverer)
	router.Use(customMiddleware.Logger())
	router.Use(customMiddleware.Authorize(repo.NewUser()))
	router.Use(customMiddleware.RoomUserCountLoader(repo.NewRoom()))
	router.Use(customMiddleware.RoomMessageCountLoader(repo.NewRoom()))
	router.Use(customMiddleware.UserLoader(repo.NewUser()))

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
	rest.NewRoutes(router, svc.NewUser())

	// graphql
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	// serve image files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Conf.Port)
	log.Fatal(http.ListenAndServe(":"+config.Conf.Port, router))
}
