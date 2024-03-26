package handler

import (
	"encoding/json"
	"movie-service/internal/variables"
	"net/http"
)

// todo validation body
// status check

func HandleUploadStatusWebhook(uploadParams variables.UploadParams) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			VideoLibraryId int    `json:"VideoLibraryId"`
			VideoGuid      string `json:"VideoGuid"`
			Status         int    `json:"Status"`
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			Logger.Error(err.Error(), http.StatusBadRequest)
			return
		}
		if input.Status != 3 {
			w.WriteHeader(http.StatusOK)
			return
		}

		UploadStatusWebhook(input.VideoGuid, input.VideoLibraryId, uploadParams.TableName, uploadParams.HostName)
		w.WriteHeader(http.StatusOK)
	}
}
