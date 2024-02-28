package tests

import (
	"movie-service/aws/awsHandlers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestListObjectHandler(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "*KEY*")
	os.Setenv("AWS_SECRET_ACESS_KEY", "*KEY*")
	os.Setenv("BUCKET_NAME", "BUCKET")

	req := httptest.NewRequest("GET", "http://localhost:8080/aws/list", nil)

	w := httptest.NewRecorder()

	awsHandlers.ListObjectsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedJSON := `[{"Name":"download.png","Size":7847}]`
	if w.Body.String() != expectedJSON {
		t.Errorf("Expected body %s, got %s", expectedJSON, w.Body.String())
	}
}
