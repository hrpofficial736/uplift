package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/models"
	"github.com/hrpofficial736/uplift/server/internal/services"
)

func handleApiRoute(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	res.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodGet {
		res.WriteHeader(401)
		return
	}
	res.WriteHeader(200)
	fmt.Fprint(res, "Welcome to the API route!")
}

func processGithubUrlHandler(res http.ResponseWriter, req *http.Request) {
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

	fmt.Println(len(request.Agents))
	response, err := services.McpConnector(request.Agents, services.CallLLM, request.Prompt)
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
