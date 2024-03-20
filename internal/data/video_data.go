package data

import (
	"context"
	"fmt"
	"movie-service/internal/database"
	"movie-service/internal/logger"
	"time"
)

type VideoRepository struct {
	Db database.Service
}

type VideoRecord struct {
	ID             int
	VideoID        string
	VideoURLHLS    string
	VideoLibraryID string
}

var VideoRepo *VideoRepository

var Logger = logger.GetLogger()

func New(db database.Service) *VideoRepository {
	return &VideoRepository{Db: db}
}

func (vr *VideoRepository) SaveVideoRecordToDB(videoLibraryID int, tableName, videoID, title, hostname string) error {
	if tableName == "" || videoID == "" {
		str := "table name and video ID cannot be empty"
		Logger.Info(str)
		return fmt.Errorf(str)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	videoURLHLS := fmt.Sprintf("https://%s/%s/playlist.m3u8", hostname, videoID)
	query := fmt.Sprintf("INSERT INTO %s (video_id, video_url_hls, library_id, title) VALUES ($1, $2, $3, $4)", tableName)

	_, err := vr.Db.ExecContext(ctx, query, videoID, videoURLHLS, videoLibraryID, title)
	if err != nil {
		Logger.Error("Failed to insert video ID:", err)
		return fmt.Errorf("error to insert into db, %w", err)
	}

	return nil
}
