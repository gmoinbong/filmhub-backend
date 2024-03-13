package streamhandlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"movie-service/internal/logger"
	"net/http"
	"strconv"
)

var Logger = logger.GetLogger()

func uploadVideo(libraryID int64, accessKey, videoId string, videoFile multipart.File) error {
	baseStreamURL := "https://video.bunnycdn.com/library/%d/videos/%s"

	url := fmt.Sprintf(baseStreamURL, libraryID, videoId)
	fileContents, err := io.ReadAll(videoFile)
	if err != nil {
		Logger.Error("Failed to read file", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileContents))
	if err != nil {
		Logger.Error("Failed to make request", err.Error())
		return err
	}

	req.Header.Set("AccessKey", accessKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Error("Failed to get response", err.Error())
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		Logger.Info("Upload failed")
		return fmt.Errorf("upload failed: %d", resp.StatusCode)
	}
	return nil
}

func HandleVideoUpload(libraryId int64, accessKey string, fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			Logger.Error("Failed to get file from request", err.Error())
			http.Error(w, "Failed to get file from request", http.StatusBadRequest)
			return
		}

		defer file.Close()

		if fileHeader.Filename != fileName {
			Logger.Info("Invalid file name")
			http.Error(w, "Invalid file name", http.StatusBadRequest)
		}
		videoID, err := createVideo(strconv.FormatInt(14, int(libraryId)), accessKey, fileName)
		if err != nil {
			Logger.Error("Failed to create video", err.Error())
			http.Error(w, "Failed to create video", http.StatusInternalServerError)
		}

		err = uploadVideo(libraryId, accessKey, videoID, file)
		if err != nil {
			Logger.Error("Failed to upload video", err.Error())
			http.Error(w, "Failed to upload video", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"message": "Video uploaded successfully", "video_id": "%s"}`, videoID)
		w.Write([]byte(response))
	}
}
