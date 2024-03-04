package router

import (
	"fmt"
	"log"
	"movie-service/aws/awsHandlers"
	"movie-service/internal/database"
	"movie-service/middleware"
	"net/http"
)

func SetupAwsRoutes() {
	middleware.SetupAwsRoute("/aws/list", http.HandlerFunc(awsHandlers.ListObjectsHandler))
	middleware.SetupAwsRoute("/aws/upload", http.HandlerFunc(awsHandlers.UploadFileHandler))
	middleware.SetupAwsRoute("/aws/update", http.HandlerFunc(awsHandlers.UpdateObjectHandler))
	middleware.SetupAwsRoute("/aws/delete", http.HandlerFunc(awsHandlers.DeleteObjectHandler))
}

func SetupRoutes() {
	middleware.SetupVideoRoutes("/films/", awsHandlers.HandleVideoRequest)
	middleware.SetupVideoRoutes("/series/", awsHandlers.HandleVideoRequest)
}

func StartServer(port string) {
	dbService := database.New()
	health := dbService.Health()
	fmt.Print(health)

	SetupAwsRoutes()
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
