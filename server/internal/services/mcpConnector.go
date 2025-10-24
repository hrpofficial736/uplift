package services

import (
	"encoding/json"
	"fmt"

	"github.com/hrpofficial736/uplift/server/internal/models"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
	"github.com/hrpofficial736/uplift/server/pkg/mcp"
)

func McpConnector(agents []string, callLLM func(string) (types.Response, error), prompt string) (interface{}, error) {

	var fResult models.CheckPointResponse
	result, err := mcp.CheckPoint(prompt, agents, callLLM)

	cleaned := utils.CleanLLMOutput(result.Text)

	if err != nil {
		return nil, fmt.Errorf("error occured on checkpoint: %s\n", err)
	}

	if err := json.Unmarshal([]byte(cleaned), &fResult); err != nil {
		return nil, fmt.Errorf("error occured on checkpoint: %s\n", err)
	}

	if fResult.Valid {
		fmt.Println("Agents to run: ", fResult.Agents)
		newAc := mcp.NewAgentCoordinator()
		fmt.Println("connector connected the handler to the mcp agent coordinator.")
		response, repoInfo, err := newAc.AddAgentAndGetAgentResponse(agents, callLLM, prompt, fResult.Url)
		if err != nil {
			return nil, fmt.Errorf("error in mcp connector while calling agent coordinator: %s", err)
		}

		return &models.Response{
			Status:   200,
			Message:  fResult.Message,
			Data:     response,
			Reviewed: true,
			RepoInfo: repoInfo,
		}, nil
	} else {
		fmt.Println("no need for mcp!")
		return &models.Response{
			Status:   200,
			Message:  fResult.Message,
			Data:     nil,
			Reviewed: false,
		}, nil
	}
}
