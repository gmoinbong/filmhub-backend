package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Unable to load AWS config:", err)
		return
	}

	client := s3.NewFromConfig(cfg)

	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("myownbucket14"),
	})
	if err != nil {
		fmt.Println("Unable to list objects in bucket:", err)
		return
	}

	fmt.Println("Objects in bucket:")
	for _, item := range resp.Contents {
		log.Printf("Name=%s Size=%d", aws.ToString(item.Key), item.Size)
	}
}
