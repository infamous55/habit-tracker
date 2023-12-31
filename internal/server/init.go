package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/infamous55/habit-tracker/internal/auth"
	"github.com/infamous55/habit-tracker/internal/ctxbridge"
	"github.com/infamous55/habit-tracker/internal/graphql"
	"github.com/infamous55/habit-tracker/internal/mongodb"
	"github.com/infamous55/habit-tracker/internal/validator"
)

func setupIndexes(db mongodb.DatabaseWrapper) error {
	indexes := []struct {
		collection string
		field      string
		unique     bool
	}{
		{"users", "email", true},
		{"groups", "user_id", false},
		{"habits", "user_id", false},
		{"habits", "group_id", false},
	}

	for _, index := range indexes {
		err := db.CreateIndex(index.collection, index.field, index.unique)
		if err != nil {
			return err
		}
	}
	return nil
}

const defaultPort = "8080"

func Init() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPost},
	}))

	// load environment variables from .env file with godotenv
	// err := godotenv.Load()
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	validator.Init()

	db := mongodb.Connect()
	defer db.Disconnect()

	err := setupIndexes(db)
	if err != nil {
		e.Logger.Fatal(err)
	}

	c := graphql.Config{Resolvers: &graphql.Resolver{
		Database: db,
	}}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(c))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.Use(extension.FixedComplexityLimit(5))

	// conditionally disable introspection in production
	if os.Getenv("ENVIRONMENT") == "production" {
		e.Use(extractPlaygroundPassword)
		srv.AroundOperations(verifyPlaygroundPassword)
	}

	e.Use(ctxbridge.EchoContextToContext)
	e.Use(auth.ExtractUserMiddleware(db))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	e.POST("/query", echo.WrapHandler(srv))

	go func() {
		if err := e.Start(":" + port); err != nil && errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
