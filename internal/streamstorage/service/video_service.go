package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"movie-service/internal/data"
	"movie-service/internal/logger"
)

var Logger = logger.GetLogger()

type VideoService struct {
	DataRepo *data.VideoRepository
}

func NewVideoService(dataRepo *data.VideoRepository) *VideoService {
	return &VideoService{DataRepo: dataRepo}
}

func (vs *VideoService) PrepareUploadVideo(libraryID, videoId string, videoFile multipart.File) (string, []byte, error) {
	baseStreamURL := "https://video.bunnycdn.com/library/%s/videos/%s"
	url := fmt.Sprintf(baseStreamURL, libraryID, videoId)

	fileContents, err := io.ReadAll(videoFile)
	if err != nil {
		Logger.Error("Failed to read file", err.Error())
		return "", nil, fmt.Errorf("failed to read file :%v", err)
	}

	return url, fileContents, nil
}
