package awsHandlers

import (
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	bucketName := os.Getenv("BUCKET_NAME")

	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "unable to get file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileUrl, err := awsOperations.UploadFileToBucket(bucketName, fileHandler)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to upload file to S3: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File upload succesfully. URL: %s", fileUrl)
}
