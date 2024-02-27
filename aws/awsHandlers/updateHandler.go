package awsHandlers

import (
	"encoding/json"
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
)

func UpdateObjectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	objectKey := r.Form.Get("objectKey")
	filepath := r.Form.Get("filePath")

	if objectKey == "" || filepath == "" {
		http.Error(w, "objectKey and filepath is now empty, but required", http.StatusBadRequest)
		return
	}

	err := awsOperations.UpdateObject(objectKey, filepath)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to update object in bucket: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Object updeted successfully"})
}
