package awsoperations

import (
	"context"
	"fmt"
	"log"
	"movie-service/aws/awsconfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListObjectsInBucket(bucketName string) {
	client, err := awsconfig.NewS3Client()
	if err != nil {
		fmt.Println("Unable to initialize S3 client:", err)
	}

	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Println("Unable to list objects in bucket:", err)
		return
	}

	fmt.Println("Objects in bucket:")
	for _, item := range resp.Contents {
		log.Printf("Name=%s Size=%d", aws.ToString(item.Key), item.Size)
	}

}
