package awsConfig

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func NewS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Unable to load AWS config:", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil

}
