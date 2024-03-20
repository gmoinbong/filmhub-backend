package handler

import (
	"movie-service/internal/data"
	"movie-service/internal/variables"
	"strconv"
)

func UploadStatusWebhook(params variables.UploadParams, videoID string, title string) error {
	videoLibraryID, err := strconv.Atoi(params.LibraryID)
	if err != nil {
		Logger.Error("Falied convert library ID", err)
		return err
	}

	err = data.VideoRepo.SaveVideoRecordToDB(videoLibraryID, params.TableName, videoID, title, params.HostName)
	if err != nil {
		Logger.Error("Falied to save vide data to database", err)
		return err
	}

	return nil
}
