package router

import (
	"movie-service/internal/cors"
	"movie-service/internal/streamstorage/handler"
	"movie-service/internal/variables"
	"net/http"
	"os"
)

var url = os.Getenv("HLS_URL")

func SetupBunnySeriesRoutes(r *Router) {
	http.Handle("/video", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleGetVideo(url))))
	http.Handle("/series", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleListVideo(variables.ListSeriesParams))))
	http.Handle("/upload/series", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleVideoUpload(variables.UploadSeriesParams))))
	http.Handle("/webhook/upload-status-series", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleUploadStatusWebhook(variables.UploadSeriesParams))))

}
func SetupBunnyFilmsRoutes(r *Router) {
	http.Handle("/films", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleListVideo(variables.ListFilmsParams))))
	http.Handle("/upload/films", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleVideoUpload(variables.UploadFilmsParams))))
	http.Handle("/webhook/upload-status-films", cors.CorsMiddleWare(http.HandlerFunc(handler.HandleUploadStatusWebhook(variables.UploadFilmsParams))))

}
