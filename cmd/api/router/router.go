package router

import (
	"fmt"
	"log/slog"
	"movie-service/internal/config"
	"movie-service/internal/cors"
	"movie-service/internal/database"
	"movie-service/internal/env"
	"movie-service/internal/logger"
	"net/http"
	"os"
	"time"
)

var Logger = logger.GetLogger()

type Router struct {
	http.ServeMux
}

func NewRouter() *Router {
	return &Router{*http.NewServeMux()}
}

func (r *Router) StartServer() {
	env.LoadEnv()
	cfg := config.New()

	dbService := database.New()
	health := dbService.Health()
	Logger.Info(fmt.Sprintf("%v", health))

	SetupBunnyRoutes(r)

	handler := setupCors(r)

	http.Handle("/", handler)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      nil,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(Logger.Handler(), slog.LevelError),
	}

	Logger.Info("starting server", "addr", srv.Addr, "env", cfg.Env)
	err := srv.ListenAndServe()
	Logger.Error("", err.Error())
	os.Exit(1)
}

func setupCors(handler http.Handler) http.Handler {
	return cors.CorsMiddleWare(handler)
}
