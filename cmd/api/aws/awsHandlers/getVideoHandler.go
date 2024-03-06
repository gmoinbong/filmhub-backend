package awsHandlers

import (
	"fmt"
	"io"
	"movie-service/cmd/api/aws/awsOperations"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func streamVideo(w http.ResponseWriter, r *http.Request, folderName string) {
	region := os.Getenv("REGION")
	bucketName := os.Getenv("BUCKET_NAME")
	videoName := r.URL.Path[len(folderName):]

	videoURL, err := awsOperations.GetVideoURL(bucketName, region, folderName, videoName)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get video URL: %v", err), http.StatusInternalServerError)
		return
	}

	resp, err := http.Head(videoURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get video information: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	videoLengthStr := resp.Header.Get("Content-Length")
	videoLength, err := strconv.ParseInt(videoLengthStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid video length: %v", err), http.StatusInternalServerError)
		return
	}

	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" {
		http.ServeFile(w, r, "video.mp4")
		return
	}

	start, end, err := parseRangeHeader(rangeHeader, videoLength)
	if err != nil {
		http.Error(w, "invalid range header", http.StatusRequestedRangeNotSatisfiable)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", strconv.FormatInt(end-start+1, 10))
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, videoLength))
	w.Header().Set("Accept-Ranges", "bytes")
	w.WriteHeader(http.StatusPartialContent)

	videoResp, err := http.Get(videoURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to open video link: %v", err), http.StatusInternalServerError)
		return
	}
	defer videoResp.Body.Close()

	buf := make([]byte, 32*1024)
	for {
		n, err := videoResp.Body.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, fmt.Sprintf("error reading video: %v", err), http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}

		if _, err := w.Write(buf[:n]); err != nil {
			http.Error(w, fmt.Sprintf("error writing response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func HandleFilmRequest(w http.ResponseWriter, r *http.Request) {
	streamVideo(w, r, "/films")
}

func HandleSerieRequest(w http.ResponseWriter, r *http.Request) {
	streamVideo(w, r, "/series")
}

func parseRangeHeader(header string, videoLength int64) (start, end int64, err error) {
	parts := strings.Split(header, "=")
	if len(parts) != 2 || parts[0] != "bytes" {
		return 0, 0, fmt.Errorf("invalid range header")
	}

	byteRange := strings.Split(parts[1], "-")
	if len(byteRange) != 2 {
		return 0, 0, fmt.Errorf("invalid byte range")
	}

	start, err = strconv.ParseInt(byteRange[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start byte")
	}

	if byteRange[1] == "" {
		end = videoLength - 1
	} else {
		end, err = strconv.ParseInt(byteRange[1], 10, 64)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid end byte")
		}
	}

	if start < 0 || start >= videoLength || end < start || end >= videoLength {
		return 0, 0, fmt.Errorf("invalid byte range values")
	}

	return start, end, nil
}
