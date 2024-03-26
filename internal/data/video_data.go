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
	VideoLibraryID int
}

var VideoRepo *VideoRepository

var Logger = logger.GetLogger()

func New(db database.Service) *VideoRepository {
	return &VideoRepository{Db: db}
}

func (vr *VideoRepository) SaveVideoRecordToDB(videoLibraryID int, videoID, tableName, hostname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	videoURLHLS := fmt.Sprintf("https://%s/%s/playlist.m3u8", hostname, videoID)

	query := fmt.Sprintf("INSERT INTO %s (video_id, library_id, video_url_hls) VALUES ($1, $2, $3)", tableName)
	_, err := vr.Db.ExecContext(ctx, query, videoID, videoLibraryID, videoURLHLS)
	if err != nil {
		Logger.Error("Failed to insert video ID:", err)
		return fmt.Errorf("error to insert into db, %w", err)
	}

	return nil
}

func (vr *VideoRepository) 	InsertVideoTitleToDB(title, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1)", tableName)
	_, err := vr.Db.ExecContext(ctx, query, title)
	if err != nil {
		Logger.Error("Failed to insert video ID:", err)
		return fmt.Errorf("error to insert into db, %w", err)
	}

	return nil
}
