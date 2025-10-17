package agents

import (
	"context"
	"fmt"

	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func MaintainabilityCritic(client *mcpclient.AgentMCPClient, server *mcpserver.AgentMCPServer, owner string, repo string, callLLM func(string) (types.Response, error), ctx context.Context, responses []interface{}) (interface{}, error) {
	fmt.Println("in maintainability agent...")
	client.Initialize()
	request := map[string]interface{}{
		"jSONRPC": "2.0",
		"method":  "maintainability",
		"params": map[string]string{
			"owner": owner,
			"repo":  repo,
		},
	}
	sentErr := client.Transport.Send(ctx, request)
	if sentErr != nil {
		fmt.Println("error in maintainability agent while sending info to mcp server")
		return nil, fmt.Errorf("error while sending the info to mcp server: %s", sentErr)
	}
	response, err := client.Transport.Receive(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in maintainability critic while receiving from the server: %s", err)
	}

	llmResponse, err := callLLM(utils.GetMaintainabilitySystemPrompt(owner, repo, response))

	if err != nil {
		return nil, fmt.Errorf("error from the llm after passing it the maintainability data: %s", err)
	}
	return types.AgentResponse{
		Data:  llmResponse,
		Agent: "maintainability",
	}, nil
}
