package utils

import (
	"path/filepath"
	"strings"
)

func ExtractTitleFromFileName(fileName string) string {
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))

	title := strings.ReplaceAll(fileNameWithoutExt, "_", " ")

	return title
}
