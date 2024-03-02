package awsHandlers

import (
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
	"os"
)

func GetVideoHandler(w http.ResponseWriter, r *http.Request) {
	region := os.Getenv("REGION")
	bucketName := os.Getenv("BUCKET_NAME")
	videoName := r.URL.Path[len("/films/"):]

	videoURL, err := awsOperations.GetVideoURL(bucketName, region, videoName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get video URL: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, videoURL, http.StatusFound)

}
