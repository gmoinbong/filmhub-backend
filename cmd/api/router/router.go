package router

import (
	"fmt"
	"log/slog"
	"movie-service/internal/config"
	"movie-service/internal/database"
	"movie-service/internal/env"
	"movie-service/internal/logger"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

type Router struct {
	http.ServeMux
}

func NewRouter() *Router {
	return &Router{*http.NewServeMux()}
}

func (r *Router) StartServer() {
	env.LoadEnv()
	logger := logger.New()
	cfg := config.New()

	dbService := database.New()
	health := dbService.Health()
	logger.Logger.Info(fmt.Sprintf("%v", health))

	handler := setupCors(r)

	SetupAwsRoutes(r)
	SetupRoutes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Logger.Handler(), slog.LevelError),
	}

	logger.Logger.Info("starting server", "addr", srv.Addr, "env", cfg.Env)
	err := srv.ListenAndServe()
	logger.Logger.Error(err.Error())
	os.Exit(1)
}

func setupCors(handler http.Handler) http.Handler {
	c := cors.Default()
	return c.Handler(handler)
}
