package streamhandlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-service/internal/utils"
	"net/http"
)

const baseStreamURL = "https://video.bunnycdn.com/library/%s/videos"

type CreateVideoRequset struct {
	Title string `json:"title"`
}

type CreateVideoResponse struct {
	VideoID string `json:"guid"`
}

func createVideo(libraryId, accessKey string, fileName string) (string, error) {

	url := fmt.Sprintf(baseStreamURL, libraryId)

	createVideoReq := CreateVideoRequset{
		Title: fileName,
	}
	createVideoReqJSON, err := json.Marshal(createVideoReq)
	if err != nil {
		Logger.Info("Failed to marshal requst JSON", err.Error())
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(createVideoReqJSON))
	if err != nil {
		Logger.Info("Failed to create video request", err.Error())
		return "", err
	}
	req.Header.Set("AccessKey", accessKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Info("Failed to get http response", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	var createVideoResp CreateVideoResponse

	err = json.NewDecoder(resp.Body).Decode(&createVideoResp)
	if err != nil {
		Logger.Info("Failed to decode response JSON", err.Error())
		return "", err
	}
	return createVideoResp.VideoID, nil
}

func HandleVideoCreate(accessKey, libId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			Logger.Error("Failed to get file from request", err)
			http.Error(w, "Failed to get file from request", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fileName := fileHeader.Filename
		title := utils.ExtractTitleFromFileName(fileName)
		Logger.Info(title)

		videoLibraryId, err := createVideo(libId, accessKey, title)
		if err != nil {
			Logger.Error("Failed to create video", err)
			http.Error(w, "Failed to create video", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		response := fmt.Sprintf(`{"message": "Video uploaded successfully", "video_id": "%s"}`, videoLibraryId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
