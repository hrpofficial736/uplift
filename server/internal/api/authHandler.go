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

func getAuthRouteHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("in auth handler")
		if req.Method != http.MethodPost {
			http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var requestBody map[string]string
		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			http.Error(res, fmt.Sprintf("invalid request body: %v", err), http.StatusBadRequest)
			return
		}

		email := requestBody["email"]
		fmt.Println(email)
		if email == "" {
			http.Error(res, "missing email", http.StatusBadRequest)
			return
		}

		rows, err := database.QueryDatabase(context.Background(), pool,
			`SELECT id, name, email, plan, prompts, plan_upgraded_at, created_at FROM "Users" WHERE email = $1`, email)
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
			fmt.Println("yes users length is zero")
			rows, err := database.QueryDatabase(context.Background(), pool,
				`INSERT INTO "Users" (name, email) VALUES ($1, $2)`,
				requestBody["name"], requestBody["email"])
			if err != nil {
				http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			return
		}
		finalResponse := UpdateResponse{
			Status:  200,
			Message: "User created successfully!",
			Data:    users[0],
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(finalResponse)
	}
}
