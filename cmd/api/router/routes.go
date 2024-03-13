package router

import (
	"movie-service/internal/cors"
	streamhandlers "movie-service/internal/streamstorage/streamHandlers"

	"net/http"
	"os"
)

var accessKeySeries = os.Getenv("BUNNY_ACCESS_KEY_SERIES")
var accessKeyFilms = os.Getenv("BUNNY_ACCESS_KEY_FILMS")
var libIdSeries = os.Getenv("BUNNY_VIDEO_LIBRARY_ID_SERIES")
var libIdFilms = os.Getenv("BUNNY_VIDEO_LIBRARY_ID_FILMS")

func SetupBunnyRoutes(r *Router) {
	http.Handle("/series", cors.CorsMiddleWare(http.HandlerFunc(streamhandlers.HandleListVideo(libIdSeries, accessKeySeries))))
	http.Handle("/films", cors.CorsMiddleWare(http.HandlerFunc(streamhandlers.HandleListVideo(libIdFilms, accessKeyFilms))))
	http.Handle("/create/series", cors.CorsMiddleWare(http.HandlerFunc(streamhandlers.HandleVideoCreate(accessKeySeries, libIdSeries))))
	http.Handle("/create/films", cors.CorsMiddleWare(http.HandlerFunc(streamhandlers.HandleVideoCreate(accessKeyFilms, libIdFilms))))

}
