package awsOperations

import (
	"fmt"
)

func GetVideoURL(bucketName, region, folderName, videoName string) (string, error) {

	videoURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s", bucketName, region, folderName, videoName)

	return videoURL, nil
}
