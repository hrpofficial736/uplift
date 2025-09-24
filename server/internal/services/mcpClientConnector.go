package services

import (
	"fmt"
	"github.com/hrpofficial736/uplift/server/pkg/mcp"
)

func McpClientConnector (agentType string, callLLM func (string) (string, error), prompt string) {
	newAc := mcp.NewAgentCoordinator()

	newAc.AddAgent(agentType, callLLM, prompt)

	fmt.Println("connector connected the handler to the mcp agent coordinator.")
}