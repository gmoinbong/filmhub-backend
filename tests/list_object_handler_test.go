package tests

import (
	"encoding/json"
	"movie-service/aws/awsHandlers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestListObjectHandler(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIASZLOJMQHRUDIXYAN")
	os.Setenv("AWS_SECRET_ACESS_KEY", "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")
	os.Setenv("BUCKET_NAME", "myownbucket14")

	req := httptest.NewRequest("GET", "http://localhost:8080/aws/list", nil)

	w := httptest.NewRecorder()

	awsHandlers.ListObjectsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedKeys := []string{"Name", "Size"}

	var recievedData []map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &recievedData); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	for _, obj := range recievedData {
		for _, key := range expectedKeys {
			if _, exists := obj[key]; !exists {
				t.Errorf("Expected key %q not found in object %+v", key, obj)
			}
		}
	}

}
