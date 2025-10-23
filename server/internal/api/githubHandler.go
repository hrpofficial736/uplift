package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/models"
	"github.com/hrpofficial736/uplift/server/internal/services"
	"github.com/hrpofficial736/uplift/server/internal/services/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func processGithubUrlHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method)
		if req.Method != http.MethodPost {
			res.WriteHeader(401)
			return
		}

		fmt.Println("in the github handler...")
		var request models.Request

		if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
			http.Error(res, "Invalid request payload", http.StatusBadRequest)
			return
		}
		promptRows, err := database.QueryDatabase(context.Background(), pool, `SELECT plan, prompts FROM "Users" WHERE email = $1`, request.Email)
		if err != nil {
			http.Error(res, "database query error", http.StatusInternalServerError)
			return
		}
		defer promptRows.Close()
		var plan string
		var prompts int
		for promptRows.Next() {
			err := promptRows.Scan(&plan, &prompts)
			if err != nil {
				http.Error(res, "error while scanning", http.StatusInternalServerError)
				return
			}
		}

		if plan == "FREE" && prompts == 3 {
			response := models.Response{
				Status:  401,
				Message: "Prompts limit Reached!",
			}
			res.WriteHeader(401)
			json.NewEncoder(res).Encode(response)
			return
		} else if plan == "PRO" && prompts == 10 {
			response := models.Response{
				Status:  401,
				Message: "Prompts limit Reached!",
			}
			res.WriteHeader(401)
			json.NewEncoder(res).Encode(response)
			return
		}

		rows, err := database.QueryDatabase(context.Background(), pool, `UPDATE "Users" SET prompts = prompts + 1 WHERE email = $1`, request.Email)
		if err != nil {
			http.Error(res, "database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		agents := []string{"security", "maintainability", "quality", "mentor"}
		response, err := services.McpConnector(agents, services.CallLLM, request.Prompt)
		if err != nil {
			serverResponseWithErr := &models.Response{
				Message: err.Error(),
			}
			json.NewEncoder(res).Encode(serverResponseWithErr)
			return
		}
		res.WriteHeader(200)
		fmt.Println("server response before sending to the client")
		json.NewEncoder(res).Encode(response)
	}
}
