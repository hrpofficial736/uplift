package api

import "net/http"

func RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("/api/", handleApiRoute)
	mux.HandleFunc("/api/github/", processGithubUrlHandler)
}
