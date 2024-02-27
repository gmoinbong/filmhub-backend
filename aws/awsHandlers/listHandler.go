package awsHandlers

import (
	"encoding/json"
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
	"os"
)

func ListObjectsHandler(w http.ResponseWriter, r *http.Request) {
	bucketName := os.Getenv("BUCKET_NAME")
	objects, err := awsOperations.ListObjectsInBucket(bucketName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to list objects in bucket: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}
