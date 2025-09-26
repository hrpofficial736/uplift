package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hrpofficial736/uplift/server/internal/utils"
	"github.com/hrpofficial736/uplift/server/pkg/mcp"
)

func McpConnector (agents []string, callLLM func (string) (utils.Response, error), prompt string) (interface{}, error) {
	
	var fResult mcp.CheckPointResponse
	result, err := mcp.CheckPoint(prompt, agents, callLLM)

	if err != nil {
		log.Fatalf("error occured on checkpoint: %s", err)
	}

	if err := json.Unmarshal([]byte(result.Text), &fResult); err != nil {
		log.Fatalf("error occured on checkpoint as conversion to the response struct failed: %s", err)
	}
	

	if fResult.Valid {
		fmt.Println("Agents to run: ", fResult.Agents)
		newAc := mcp.NewAgentCoordinator()
		fmt.Println("connector connected the handler to the mcp agent coordinator.")
		return newAc.AddAgent(agents, callLLM, prompt)
		
	} else {
		return nil, fmt.Errorf("checkpoint failed: %s", fResult.Message)
	}	
}