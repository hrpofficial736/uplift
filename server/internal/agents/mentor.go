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

	llmResponse, err := callLLM(utils.GetSecuritySystemPrompt(owner, repo, responses))

	if err != nil {
		return nil, fmt.Errorf("error from the llm after passing it the mentor data: %s", err)
	}
	return llmResponse, nil
}
