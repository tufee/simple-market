package api

import "net/http"

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", Hello)
	mux.HandleFunc("POST /signup", Signup)

	return mux
}
