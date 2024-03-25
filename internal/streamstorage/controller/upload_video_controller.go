package controller

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"movie-service/internal/logger"
	"movie-service/internal/streamstorage/service"
	"movie-service/internal/variables"
	"net/http"
)

type UploadVideoController struct {
	VideoService *service.VideoService
}

var Logger = logger.GetLogger()

func NewVideoController(videoService service.VideoService) *UploadVideoController {
	return &UploadVideoController{VideoService: &videoService}
}

func (vc *UploadVideoController) UploadVideo(uploadParams variables.UploadParams, videoId string, file multipart.File) error {

	url, fileContents, err := vc.VideoService.PrepareUploadVideo(uploadParams.LibraryID, videoId, file)
	if err != nil {
		Logger.Error("Failed to prepare upload video", err.Error())
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileContents))
	if err != nil {
		Logger.Error("Failed to make request", err.Error())
		return err
	}

	req.Header.Set("AccessKey", uploadParams.AccessKey)
	req.Header.Set("Content-Type", "application/json")

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
