package router

import (
	"movie-service/cmd/api/aws/awsHandlers"
)

func SetupAwsRoutes(r *Router) {
	r.HandleFunc("/aws/list", awsHandlers.ListObjectsHandler)
	r.HandleFunc("/aws/upload", awsHandlers.UploadFileHandler)
	r.HandleFunc("/aws/update", awsHandlers.UpdateObjectHandler)
	r.HandleFunc("/aws/delete", awsHandlers.DeleteObjectHandler)
}

func SetupRoutes(r *Router) {
	r.HandleFunc("/films/", awsHandlers.HandleFilmRequest)
	r.HandleFunc("/series/", awsHandlers.HandleSerieRequest)
}
