package api

import "net/http"

func RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/api/github/", processGithubUrlHandler)
	mux.HandleFunc("/api/", handleApiRoute)
}
