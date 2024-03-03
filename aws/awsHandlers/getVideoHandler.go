package awsHandlers

import (
	"fmt"
	"io"
	"log"
	"movie-service/aws/awsOperations"
	"net/http"
	"os"
)

func HandleVideoRequest(w http.ResponseWriter, r *http.Request, folderName string) {
	region := os.Getenv("REGION")
	bucketName := os.Getenv("BUCKET_NAME")
	videoName := r.URL.Path[len(folderName):]

	videoURL, err := awsOperations.GetVideoURL(bucketName, region, folderName, videoName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get video URL: %v", err), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(videoURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to open link: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("error streaming vide: %v", err)
		return
	}

	http.Redirect(w, r, videoURL, http.StatusFound)

}
