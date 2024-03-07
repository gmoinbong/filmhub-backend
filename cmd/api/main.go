package main

import (
	"movie-service/cmd/api/router"
)

func main() {
	//TODO: IMDB rating parser
	r := router.NewRouter()
	r.StartServer()
}
