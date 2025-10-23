package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/models"
	"github.com/hrpofficial736/uplift/server/internal/services/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateRequest struct {
	Email   string `json:"email"`
	Plan    string `json:"plan"`
	Prompts int    `json:"prompts"`
}

type UpdateResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    models.User `json:"data"`
}

func getUserInfo(pool *pgxpool.Pool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestBody map[string]string
		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			http.Error(res, "invalid request body", http.StatusBadRequest)
			return
		}
		rows, err := database.QueryDatabase(context.Background(), pool,
			`SELECT id, name, email, plan, prompts, plan_upgraded_at, created_at FROM "Users" WHERE email = $1`, requestBody["email"])
		if err != nil {
			http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User

		for rows.Next() {
			var user models.User
			err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Plan, &user.Prompts, &user.Plan_upgraded_at, &user.CreatedAt)
			if err != nil {
				http.Error(res, fmt.Sprintf("scan error: %v", err), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		if len(users) == 0 {
			http.Error(res, "user not found", http.StatusNotFound)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(users[0])
	}
}

func updateUser(pool *pgxpool.Pool) http.HandlerFunc {
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

		rows, err := database.QueryDatabase(context.Background(), pool, `UPDATE "Users" SET prompts = $1 WHERE email = $2`, requestBody.Prompts, requestBody.Email)
		if err != nil {
			http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		res.WriteHeader(200)
		response := UpdateResponse{
			Status:  200,
			Message: "User updated successfully!",
		}

		json.NewEncoder(res).Encode(response)
	}
}
