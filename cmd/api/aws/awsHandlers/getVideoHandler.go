package awsHandlers

import (
	"context"
	"fmt"
	"io"
	"movie-service/cmd/api/aws/awsOperations"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", videoURL, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating request: %v", err), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to open video link: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	videoLengthStr := resp.Header.Get("Content-Length")
	videoLength, err := strconv.ParseInt(videoLengthStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid video length: %v", err), http.StatusInternalServerError)
		return
	}

	start, end, err := parseRangeHeader(r.Header.Get("Range"), videoLength)
	if err != nil {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		w.Write([]byte("invalid range header"))
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", strconv.FormatInt(end-start+1, 10))
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, videoLength))
	w.Header().Set("Accept-Ranges", "bytes")
	w.WriteHeader(http.StatusPartialContent)

	if start > 0 {
		if _, err := io.CopyN(io.Discard, resp.Body, start); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error skipping bytes: %v", err)))
			return
		}
	}

	if _, err := io.CopyN(w, resp.Body, end-start+1); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error sending video data: %v", err)))
		return
	}
}

func HandleFilmRequest(w http.ResponseWriter, r *http.Request) {
	streamVideo(w, r, "/films")
}

func HandleSerieRequest(w http.ResponseWriter, r *http.Request) {
	streamVideo(w, r, "/series")
}

func parseRangeHeader(header string, videoLength int64) (start, end int64, err error) {
	if header == "" {
		return 0, 0, nil
	}

	parts := strings.SplitN(header, "=", 2)
	if len(parts) != 2 || parts[0] != "bytes" {
		return 0, 0, fmt.Errorf("invalid range header")
	}

	byteRange := strings.SplitN(parts[1], "-", 2)
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
