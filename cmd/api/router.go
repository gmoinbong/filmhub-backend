package main

import (
	"fmt"
	"log/slog"
	"movie-service/cmd/api/aws/awsHandlers"
	"movie-service/internal/config"
	"movie-service/internal/database"
	"movie-service/internal/env"
	"movie-service/internal/logger"
	"movie-service/internal/middleware"
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
