package router

import (
	"log"
	"movie-service/aws/awsHandlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/aws/list", awsHandlers.ListObjectsHandler)
	http.HandleFunc("/aws/upload", awsHandlers.UploadFileHandler)
	http.HandleFunc("/aws/update", awsHandlers.UpdateObjectHandler)
	http.HandleFunc("/aws/delete", awsHandlers.DeleteObjectHandler)
}

func StartServer(port string) {
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
