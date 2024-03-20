package handler

import (
	"fmt"
	"net/http"
)

func HandleGetVideo(videoURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", videoURL)
	}
}
