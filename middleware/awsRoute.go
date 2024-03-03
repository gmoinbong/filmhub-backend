package middleware

import (
	"net/http"
)

func SetupAwsRoute(path string, handler http.Handler) {
	http.Handle(path, EnableCors(handler))
}
