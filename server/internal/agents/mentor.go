package agents

import (
	"context"
	"fmt"

	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func Mentor(client *mcpclient.AgentMCPClient, server *mcpserver.AgentMCPServer, owner string, repo string, callLLM func(string) (types.Response, error), ctx context.Context, responses []interface{}) (interface{}, error) {
	fmt.Println("in mentor agent...")

	llmResponse, err := callLLM(utils.GetMentorSystemPrompt(owner, repo, responses))
	if llmResponse.Text == "" {
		return nil, fmt.Errorf("model overloaded, please try after some time")
	}
	if err != nil {
		return nil, fmt.Errorf("error from the llm after passing it the mentor data: %s", err)
	}
	return types.AgentResponse{
		Data:  llmResponse,
		Agent: "mentor",
	}, nil
}
