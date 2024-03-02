package awsOperations

import (
	"fmt"
)

func GetVideoURL(bucketName, region, videoName string) (string, error) {

	videoURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, videoName)

	return videoURL, nil
}
