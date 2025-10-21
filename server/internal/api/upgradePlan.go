package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/models"
	"github.com/hrpofficial736/uplift/server/internal/services/database"
	"github.com/hrpofficial736/uplift/server/internal/services/payment"
	"github.com/jackc/pgx/v5/pgxpool"
)

func upgradePlan(pool *pgxpool.Pool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestBody UpdateRequest
		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			http.Error(res, fmt.Sprintf("invalid request body: %v", err), http.StatusBadRequest)
			return
		}

		if requestBody.Email == "" {
			http.Error(res, "email not found in the request body", http.StatusBadRequest)
			return
		}

		fetchUserRows, err := database.QueryDatabase(context.Background(), pool,
			`SELECT id from "Users" WHERE email = $1`, requestBody.Email)

		if err != nil {
			http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
			return
		}
		defer fetchUserRows.Close()
		var userId string
		for fetchUserRows.Next() {
			err := fetchUserRows.Scan(&userId)
			if err != nil {
				http.Error(res, "error while scanning", http.StatusInternalServerError)
				return
			}
		}
		if userId == "" {
			http.Error(res, "failed to get user id", http.StatusInternalServerError)
			return
		}

		url, err := payment.HandleCreateCheckoutSession()
		if err != nil {
			http.Error(res, fmt.Sprintf("stripe payment error: %s", err), http.StatusInternalServerError)
			return
		}

		if url == "" {
			http.Error(res, "failed to get checkout url", http.StatusInternalServerError)
			return
		}
		rows, err := database.QueryDatabase(context.Background(), pool,
			`UPDATE "Users" SET plan = $1 WHERE id = $2`, "PRO", userId)
		if err != nil {
			http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		res.WriteHeader(200)
		response := UpdateResponse{
			Status:  200,
			Message: "Plan upgraded successfully!",
			Data: models.User{
				Id: userId,
			},
		}

		json.NewEncoder(res).Encode(response)
	}
}
