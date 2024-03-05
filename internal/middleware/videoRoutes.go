package middleware

import "net/http"

func SetupVideoRoutes(prefix string, handleFunc func(http.ResponseWriter, *http.Request, string)) {
	http.Handle(prefix, EnableCors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleFunc(w, r, prefix)
	})))
}
