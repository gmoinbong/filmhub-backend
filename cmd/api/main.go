package main

import "movie-service/cmd/api/router"

func main() {
	r := router.NewRouter()
	r.StartServer()
}
