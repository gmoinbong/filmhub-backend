package variables

type ListParams struct {
	LibraryID string
	AccessKey string
}

var ListSeriesParams = ListParams{
	LibraryID: getEnv("BUNNY_VIDEO_LIBRARY_ID_SERIES"),
	AccessKey: getEnv("BUNNY_ACCESS_KEY_SERIES"),
}

var ListFilmsParams = ListParams{
	LibraryID: getEnv("BUNNY_VIDEO_LIBRARY_ID_FILMS"),
	AccessKey: getEnv("BUNNY_ACCESS_KEY_FILMS"),
}
