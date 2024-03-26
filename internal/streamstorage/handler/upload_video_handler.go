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
	uploadStatusHandler := HandleUploadStatusWebhook(uploadParams)

	return func(w http.ResponseWriter, r *http.Request) {
		Logger.Info("Incoming request:", r.Method, r.URL.Path)

		if r.Method == http.MethodPut && r.Header.Get("X-Callback-Key") == variables.ListSeriesParams.AccessKey || r.Header.Get("X-Callback-Key") == variables.ListFilmsParams.AccessKey {
			uploadStatusHandler(w, r)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			Logger.Error("Failed to get file from request", err.Error())
			http.Error(w, "Failed to get file from request", http.StatusBadRequest)
			return
		}

		defer file.Close()

		fileName := fileHeader.Filename
		title, err := utils.ExtractTitleFromFileName(fileName)
		if err != nil {
			Logger.Error("Failed to extract title from file name", err)
		}

		Logger.Info(title)

		if fileHeader.Filename != fileName {
			Logger.Info("Invalid file name")
			http.Error(w, "Invalid file name", http.StatusBadRequest)
			return
		}

		videoID, err := createVideo(uploadParams.LibraryID, uploadParams.AccessKey, title, uploadParams.TableName)
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
