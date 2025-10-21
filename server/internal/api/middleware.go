package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hrpofficial736/uplift/server/internal/config"
)

func MiddleWare(next http.Handler) http.Handler {
	client := config.ConfigLoad().ClientUrl
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", client)
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println("in middleware")
		if req.Method == http.MethodOptions {
			fmt.Println("yes it is an options request.")
			res.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(res, req)
	})
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		cfg := config.ConfigLoad()
		supabaseSecret := cfg.SupabaseJWTSecret
		authHeader := req.Header.Get("Authorization")

		if authHeader == "" {
			fmt.Println("missing auth header")
			http.Error(res, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("invalid auth header")
			http.Error(res, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		tokenFromHeader := parts[1]

		token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(supabaseSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("invalid token")
			http.Error(res, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), "claims", token.Claims)
		next(res, req.WithContext(ctx))
	}

}
