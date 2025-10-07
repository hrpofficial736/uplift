package services

import (
	"encoding/json"
	"fmt"
	"log"

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
		log.Fatalf("error occured on checkpoint: %s", err)
	}

	if err := json.Unmarshal([]byte(cleaned), &fResult); err != nil {
		log.Fatalf("error occured on checkpoint as conversion to the response struct failed: %s", err)
	}

	if fResult.Valid {
		fmt.Println("Agents to run: ", fResult.Agents)
		newAc := mcp.NewAgentCoordinator()
		fmt.Println("connector connected the handler to the mcp agent coordinator.")
		return newAc.AddAgentAndGetAgentResponse(agents, callLLM, prompt, fResult.Url)

	} else {
		fmt.Println("no need for mcp!")
		return nil, fmt.Errorf("%s", fResult.Message)
	}
}
