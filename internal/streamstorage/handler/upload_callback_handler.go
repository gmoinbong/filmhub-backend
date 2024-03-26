package handler

import (
	"movie-service/internal/data"
)

func UploadStatusWebhook(videoID string, videoLibraryID int, tableName, hostname string) error {
	err := data.VideoRepo.SaveVideoRecordToDB(videoLibraryID, videoID, tableName, hostname)
	if err != nil {
		Logger.Error("Falied to save vide data to database", err)
		return err
	}

	return nil
}
