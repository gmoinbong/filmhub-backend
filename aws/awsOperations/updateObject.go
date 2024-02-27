package awsOperations

import (
	"context"
	"fmt"
	"log"
	"movie-service/aws/awsConfig"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UpdateObject(objectKey string, filePath string) error {
	bucket := os.Getenv("BUCKET_NAME")
	client, err := awsConfig.NewS3Client()
	if err != nil {
		return fmt.Errorf("unable to initialize S3 client: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("unable to udpate object in bucket: %v", err)
	}

	log.Printf("Object %s updated successfully\n", objectKey)

	return nil
}
