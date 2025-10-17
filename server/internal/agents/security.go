package agents

import (
	"context"
	"fmt"

	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func SecurityCritic(client *mcpclient.AgentMCPClient, server *mcpserver.AgentMCPServer, owner string, repo string, callLLM func(string) (types.Response, error), ctx context.Context, responses []interface{}) (interface{}, error) {
	fmt.Println("in security agent...")
	client.Initialize()
	request := map[string]interface{}{
		"jSONRPC": "2.0",
		"method":  "security",
		"params": map[string]string{
			"owner": owner,
			"repo":  repo,
		},
	}
	sentErr := client.Transport.Send(ctx, request)
	if sentErr != nil {
		fmt.Println("error in security agent while sending info to mcp server")
		return nil, fmt.Errorf("error while sending the info to mcp server: %s", sentErr)
	}
	response, err := client.Transport.Receive(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in security critic while receiving from the server: %s", err)
	}

	llmResponse, err := callLLM(utils.GetSecuritySystemPrompt(owner, repo, response))

	if err != nil {
		return nil, fmt.Errorf("error from the llm after passing it the security data: %s", err)
	}
	return types.AgentResponse{
		Data:  llmResponse,
		Agent: "security",
	}, nil
}
