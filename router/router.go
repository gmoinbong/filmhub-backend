package router

import (
	"log"
	"movie-service/aws/awsHandlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/list-aws", awsHandlers.ListObjectsHandler)

	http.HandleFunc("/upload-aws", awsHandlers.UploadFileHandler)
}

func StartServer(port string) {
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
