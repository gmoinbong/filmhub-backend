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

func DeleteObject(objectKey string) error {
	bucketName := os.Getenv("BUCKET_NAME")

	client, err := awsConfig.NewS3Client("vladyslav")
	if err != nil {
		return fmt.Errorf("unable to initialize S3 client :%v", err)
	}

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Println("unable to delete object from bucket:", err)
	}

	log.Printf("Object %s deleted successfully\n", objectKey)
	return nil
}
