package router

import (
	"log"
	"movie-service/aws/awsHandlers"
	"net/http"
)


func SetupAwsRoutes() {
	http.HandleFunc("/aws/list", awsHandlers.ListObjectsHandler)
	http.HandleFunc("/aws/upload", awsHandlers.UploadFileHandler)
	http.HandleFunc("/aws/update", awsHandlers.UpdateObjectHandler)
	http.HandleFunc("/aws/delete", awsHandlers.DeleteObjectHandler)
}

func SetupRoutes() {
	//TODO: add profile logic for default useres and admin panel
	http.HandleFunc("/films/", func(w http.ResponseWriter, r *http.Request) {
		awsHandlers.HandleVideoRequest(w, r, "films")
	})
	http.HandleFunc("/series/", func(w http.ResponseWriter, r *http.Request) {
		awsHandlers.HandleVideoRequest(w, r, "series")
	})
}

func StartServer(port string) {
	SetupAwsRoutes()
	SetupRoutes()

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
