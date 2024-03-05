package router

import (
	"fmt"
	"log/slog"
	"movie-service/aws/awsHandlers"
	"movie-service/internal/config"
	"movie-service/internal/database"
	"movie-service/internal/logger"
	"movie-service/middleware"
	"movie-service/router/env"
	"net/http"
	"os"
	"time"
)

func setupAwsRoutes() {
	middleware.SetupAwsRoute("/aws/list", http.HandlerFunc(awsHandlers.ListObjectsHandler))
	middleware.SetupAwsRoute("/aws/upload", http.HandlerFunc(awsHandlers.UploadFileHandler))
	middleware.SetupAwsRoute("/aws/update", http.HandlerFunc(awsHandlers.UpdateObjectHandler))
	middleware.SetupAwsRoute("/aws/delete", http.HandlerFunc(awsHandlers.DeleteObjectHandler))
}

func setupRoutes() {
	middleware.SetupVideoRoutes("/films/", awsHandlers.HandleVideoRequest)
	middleware.SetupVideoRoutes("/series/", awsHandlers.HandleVideoRequest)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func StartServer() {
	env.LoadEnv()
	logger := logger.New()
	cfg := config.New()

	dbService := database.New()
	health := dbService.Health()
	logger.Logger.Info(fmt.Sprintf("%v", health))

	setupAwsRoutes()
	setupRoutes()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      http.HandlerFunc(healthCheckHandler),
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
