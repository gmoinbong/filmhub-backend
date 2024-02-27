package router

import (
	"log"
	"movie-service/aws/awsOperations"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/upload-aws", awsOperations.UploadFileHandler)

	http.HandleFunc("list-aws", awsOperations.ListObjectsHandler)
}

func StartServer(port string) {
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
