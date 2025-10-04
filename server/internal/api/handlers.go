package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/services"
)



func handleApiRoute (res http.ResponseWriter, req *http.Request) {
	if (req.Method != http.MethodGet) {
		res.WriteHeader(401);
		return;
	}
	res.WriteHeader(200);
	fmt.Fprint(res, "Welcome to the API route!");
}


type Request struct {
	Prompt string
	Agents []string
}

func processGithubUrlHandler (res http.ResponseWriter, req *http.Request) {
	if (req.Method != http.MethodPost) {
		res.WriteHeader(401);
		return;
	}
	
	var request Request;

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Invalid request payload", http.StatusBadRequest);
		return;
	}

	services.McpConnector(request.Agents, services.CallLLM, request.Prompt)	

	// responseFromGithubService := services.FetchRepoInfo(request.Url);
	// defer req.Body.Close();

	// res.Header().Set("Content-Type", "application/json");
	// response := map[string] any{
	// 	"status": "OK",
	// 	"url": request.Url,
	// 	"repo_info": responseFromGithubService,
	// }
	// res.WriteHeader(http.StatusOK);
	// json.NewEncoder(res).Encode(response);
}