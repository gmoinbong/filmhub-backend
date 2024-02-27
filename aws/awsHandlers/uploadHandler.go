package awsHandlers

import (
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "unable to get file from request", http.StatusBadRequest)
		http.Error(w, "unable to get file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileUrl, err := awsOperations.UploadFileToBucket("myownbucket14", fileHandler)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to upload file to S3: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File upload succesfully. URL: %s", fileUrl)
}
