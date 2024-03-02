package awsHandlers

import (
	"fmt"
	"movie-service/aws/awsOperations"
	"net/http"
	"os"
)

func HandleVideoRequest(w http.ResponseWriter, r *http.Request, folderName string) {
	region := os.Getenv("REGION")
	bucketName := os.Getenv("BUCKET_NAME")
	videoName := r.URL.Path[len("/"+folderName+"/"):]

	videoURL, err := awsOperations.GetVideoURL(bucketName, region, folderName, videoName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get video URL: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, videoURL, http.StatusFound)

}
