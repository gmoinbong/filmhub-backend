package awsOperations

import (
	"context"
	"fmt"
	"log"
	"movie-service/aws/awsConfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Object struct {
	Name string
	Size int64
}

func ListObjectsInBucket(bucketName string) ([]S3Object, error) {
	client, err := awsConfig.NewS3Client()
	if err != nil {
		return nil, fmt.Errorf("unable to initialize S3 client: %v", err)
	}

	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Println("Unable to list objects in bucket:", err)
		return nil, err
	}

	fmt.Println("Objects in bucket:")
	var objects []S3Object
	for _, item := range resp.Contents {
		obj := S3Object{
			Name: aws.ToString(item.Key),
			Size: *item.Size,
		}
		objects = append(objects, obj)
	}

	return objects, nil

}
