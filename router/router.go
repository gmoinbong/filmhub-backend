package router

import (
	"log"
	"movie-service/aws/awsHandlers"
	"movie-service/internal/handlers"
	"net/http"
)

func SetupAwsRoutes() {
	http.HandleFunc("/aws/list", awsHandlers.ListObjectsHandler)
	http.HandleFunc("/aws/upload", awsHandlers.UploadFileHandler)
	http.HandleFunc("/aws/update", awsHandlers.UpdateObjectHandler)
	http.HandleFunc("/aws/delete", awsHandlers.DeleteObjectHandler)
}

func SetupRoutes() {
	http.HandleFunc("/films/", awsHandlers.GetVideoHandler)
	http.HandleFunc("/series/", handlers.GetSeriesHandler)
}

func StartServer(port string) {
	SetupAwsRoutes()
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
