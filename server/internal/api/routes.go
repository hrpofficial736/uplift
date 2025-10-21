package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouter(mux *http.ServeMux, pool *pgxpool.Pool) {
	mux.HandleFunc("/api/github", AuthMiddleware(processGithubUrlHandler))
	mux.Handle("/api/update-user", AuthMiddleware(updateUser(pool)))
	mux.Handle("/api/upgrade-plan", AuthMiddleware(upgradePlan(pool)))
	mux.Handle("/api/create-checkout-session", AuthMiddleware(getCheckoutSession(pool)))
	mux.Handle("/api/auth", AuthMiddleware(getAuthRouteHandler(pool)))
	mux.HandleFunc("/api", AuthMiddleware(processGithubUrlHandler))
}
