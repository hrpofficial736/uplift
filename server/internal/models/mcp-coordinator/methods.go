package mcpcoordinator

import (
	"context"
	"fmt"

	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func (ac *AgentCoordinator) AddAgent(agents []string, callLLM func(string) (types.Response, error), prompt string, url string) (interface{}, error) {

	for _, agentType := range agents {
		transport := ac.TransportManager.CreateTransport(agentType)
		ctx, cancel := context.WithCancel(context.Background())

		client := &mcpclient.AgentMCPClient{
			AgentType: agentType,
			CallLLM:   callLLM,
			Transport: transport,
			Ctx:       ctx,
			Cancel:    cancel,
		}

		client.Initialize()

		server := &mcpserver.AgentMCPServer{
			ServerId:  agentType,
			Transport: transport,
		}

		ac.McpClients[agentType] = client
		ac.McpServers[agentType] = server

		utils.AgentTypeToFunctionMapping[agentType](client, server, url, callLLM)
	}

	return nil, fmt.Errorf("error aa rha h")
}
