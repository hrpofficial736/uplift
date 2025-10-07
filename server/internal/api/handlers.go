package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/services"
)

func handleApiRoute(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(401)
		return
	}
	res.WriteHeader(200)
	fmt.Fprint(res, "Welcome to the API route!")
}

type Request struct {
	Prompt string
	Agents []string
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func processGithubUrlHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(401)
		return
	}
	res.Header().Set("Content-Type", "application/json")

	fmt.Println("in the github handler...")
	var request Request

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := services.McpConnector(request.Agents, services.CallLLM, request.Prompt)
	if err != nil {
		serverResponse := &Response{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(serverResponse)
		return
	}
	fmt.Println(response)
	serverResponse := &Response{
		Status:  200,
		Message: "Fetched response successfully!",
		Data:    response,
	}
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(serverResponse)
}
