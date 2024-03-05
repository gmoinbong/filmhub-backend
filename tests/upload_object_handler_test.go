package tests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"movie-service/cmd/api/aws/awsHandlers"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadObjectHandler(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIASZLOJMQHRUDIXYAN")
	os.Setenv("AWS_SECRET_ACESS_KEY", "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")
	os.Setenv("BUCKET_NAME", "myownbucket14")

	file, err := os.Open("/home/vladyslav/Downloads/download.png")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("Failed to copy file contents: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "http://localhost:8080/aws/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	awsHandlers.UploadFileHandler(w, req)

	fmt.Println("Response body:", w.Body.String())

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}
