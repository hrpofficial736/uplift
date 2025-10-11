package api

import (
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
)

func MiddleWare(next http.Handler) http.Handler {
	client := config.ConfigLoad().ClientUrl
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", client)
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		res.Header().Set("Access-Control-Allow-Credentials", "true")

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(res, req)
	})
}
