package adminpanel

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetFileLink(cfg aws.Config, key string) (string, error) {
	bucketName := os.Getenv("BUCKET_NAME")

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	presignUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key)},

		s3.WithPresignExpires(time.Minute*15))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(presignUrl.URL)
	return presignUrl.URL, nil
}
