package awsHandlers

import (
	"fmt"
	"movie-service/internal/aws/awsOperations"
	"net/http"
)

func DeleteObjectHandler(w http.ResponseWriter, r *http.Request) {
	objectKey := r.URL.Query().Get("objectKey")
	if objectKey == "" {
		http.Error(w, "objectKey parametr now empty, but is required", http.StatusBadRequest)
		return
	}
	err := awsOperations.DeleteObject(objectKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to delete object from bucket: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
