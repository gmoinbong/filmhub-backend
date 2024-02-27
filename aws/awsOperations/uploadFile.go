package awsOperations

import (
	"context"
	"fmt"
	"mime/multipart"
	"movie-service/aws/awsConfig"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func uploadFileToBucket(bucketName string, fileHeader *multipart.FileHeader) (string, error) {
	client, err := awsConfig.NewS3Client()
	if err != nil {
		return "", fmt.Errorf("unable to use client")
	}
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("unable to open file")
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("unable to upload file to S3: %v", err)
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName), nil
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "unable to get file from request", http.StatusBadRequest)
		http.Error(w, "unable to get file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileUrl, err := uploadFileToBucket("myownbucket14", fileHandler)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to upload file to S3: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File upload succesfully. URL: %s", fileUrl)
}
