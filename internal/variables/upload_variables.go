package variables

import "movie-service/internal/logger"

var Logger = logger.GetLogger()

type UploadParams struct {
	LibraryID string
	AccessKey string
	TableName string
	HostName  string
}

var UploadSeriesParams = UploadParams{
	LibraryID: getEnv("BUNNY_VIDEO_LIBRARY_ID_SERIES"),
	AccessKey: getEnv("BUNNY_ACCESS_KEY_SERIES"),
	TableName: "series",
	HostName:  getEnv("BUNNY_CDN_HOSTNAME_SERIES"),
}

var UploadFilmsParams = UploadParams{
	LibraryID: getEnv("BUNNY_VIDEO_LIBRARY_ID_FILMS"),
	AccessKey: getEnv("BUNNY_ACCESS_KEY_FILMS"),
	TableName: "films",
	HostName:  getEnv("BUNNY_CDN_HOSTNAME_FILMS"),
}
