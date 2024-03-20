package handler

import (
	"fmt"
	"movie-service/internal/data"
	"movie-service/internal/logger"
	"movie-service/internal/streamstorage/controller"
	"movie-service/internal/streamstorage/service"
	"movie-service/internal/utils"
	"movie-service/internal/variables"
	"net/http"
)

var Logger = logger.GetLogger()

func HandleVideoUpload(uploadParams variables.UploadParams) http.HandlerFunc {
	controller := controller.NewVideoController(*service.NewVideoService(data.VideoRepo))
	return func(w http.ResponseWriter, r *http.Request) {
		Logger.Info("Incoming request:", r.Method, r.URL.Path)

		if r.Method == http.MethodPost {
			if r.Header.Get("X-Callback-Key") == variables.ListSeriesParams.AccessKey {
				err := r.ParseForm()
				if err != nil {
					http.Error(w, "Failed to parse request body", http.StatusBadRequest)
					return
				}

				videoID := r.Form.Get("VideoGuid")
				title := r.Form.Get("title")
				if videoID == "" || title == "" {
					http.Error(w, "Video ID or title is empty", http.StatusBadRequest)
					return
				}
				Logger.Info("Received video ID:", videoID, " and title:", title)

				err = UploadStatusWebhook(uploadParams, videoID, title)
				if err != nil {
					Logger.Error("Failed to handle upload status webhook", err.Error())
					http.Error(w, "Failed to handle upload status webhook", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				return
			}
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			Logger.Error("Failed to get file from request", err.Error())
			http.Error(w, "Failed to get file from request", http.StatusBadRequest)
			return
		}

		defer file.Close()

		fileName := fileHeader.Filename
		title := utils.ExtractTitleFromFileName(fileName)
		Logger.Info(title)

		if fileHeader.Filename != fileName {
			Logger.Info("Invalid file name")
			http.Error(w, "Invalid file name", http.StatusBadRequest)
			return
		}

		videoID, err := createVideo(uploadParams.LibraryID, uploadParams.AccessKey, title)
		if err != nil {
			Logger.Error("Failed to create video", err.Error())
			http.Error(w, "Failed to create video", http.StatusInternalServerError)
			return
		}

		err = controller.UploadVideo(uploadParams, videoID, file)
		if err != nil {
			Logger.Error("Failed to upload video", err.Error())
			http.Error(w, "Failed to upload video", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := fmt.Sprintf(`{"request was": "true",  "video_id": "%s"}`, videoID)
		w.Write([]byte(resp))
	}
}
