package main

import (
	"movie-service/router"
	"movie-service/router/env"
)

func main() {
	env.LoadEnv()
	port := env.GoPort()

	router.StartServer(port)
}
