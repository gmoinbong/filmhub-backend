package tests

import (
	"movie-service/internal/aws/awsHandlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const expectedContentType string = "video/mp4"

func TestHandleFilmRequset(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/films/Klan_Soprano_1_sezon_-_1_seriya.mp4", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(awsHandlers.HandleFilmRequest)

	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusPartialContent {
		t.Errorf("handler returned wrong status code, got: %v expected: %v",
			status, http.StatusPartialContent)
	}

	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected type content, got: %v, expected: %v",
			contentType, expectedContentType)
	}

}
