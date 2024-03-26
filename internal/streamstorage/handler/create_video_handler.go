package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-service/internal/data"
	"net/http"
)

const baseStreamURL = "https://video.bunnycdn.com/library/%s/videos"

type CreateVideoRequset struct {
	Title string `json:"title"`
}

type CreateVideoResponse struct {
	VideoID string `json:"guid"`
}

func createVideo(libraryID, accessKey string, fileName, tableName string) (string, error) {

	url := fmt.Sprintf(baseStreamURL, libraryID)

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

	err = data.VideoRepo.InsertVideoTitleToDB(fileName, tableName)
	if err != nil {
		Logger.Error("Falied to save vide data to database", err)
		return "", err
	}

	err = json.NewDecoder(resp.Body).Decode(&createVideoResp)
	if err != nil {
		Logger.Info("Failed to decode response JSON", err.Error())
		return "", err
	}

	return createVideoResp.VideoID, nil
}
