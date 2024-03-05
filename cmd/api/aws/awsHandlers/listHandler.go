package awsHandlers

import (
	"encoding/json"
	"fmt"
	"movie-service/cmd/api/aws/awsOperations"
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

	jsonData, err := json.Marshal(objects)
	if err != nil {
		http.Error(w, "unable to marshal objects to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
