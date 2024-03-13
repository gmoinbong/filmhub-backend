package streamhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VideoListResponse struct {
	Items []Video `json:"items"`
}

type Video struct {
	VideoLibraryId int64  `json:"videoLibraryId"`
	Guid           string `json:"guid,omitempty"`
	Title          string `json:"title,omitempty"`
	Length         int32  `json:"length"`
	Status         int    `json:"status"`
	Height         int    `json:"height"`
}

var baseURL = "https://video.bunnycdn.com/library/%s/videos"

func listStreamVideos(libraryId, accessKey string) (VideoListResponse, error) {
	url := fmt.Sprintf(baseURL, libraryId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Logger.Error(err.Error())
		return VideoListResponse{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("AccessKey", accessKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Logger.Error(err.Error())
		return VideoListResponse{}, err
	}
	defer res.Body.Close()

	var response VideoListResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		Logger.Error(err.Error())
		return VideoListResponse{}, err
	}

	return response, nil
}

func HandleListVideo(libraryId, accesKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		objects, err := listStreamVideos(libraryId, accesKey)
		if err != nil {
			Logger.Error(err.Error())
			http.Error(w, "Failed to list object from Bunny Stream Storage", http.StatusInternalServerError)
		}

		jsonData, err := json.Marshal(objects)
		if err != nil {
			Logger.Error(err.Error())
			http.Error(w, "Failed to encode response to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonData)
		if err != nil {
			Logger.Error(err.Error())
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}

}
