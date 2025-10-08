package mcpcoordinator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hrpofficial736/uplift/server/internal/constants"
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/github"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

func (ac *AgentCoordinator) AddAgentAndGetAgentResponse(agents []string, callLLM func(string) (types.Response, error), prompt string, url string) ([]interface{}, error) {
	info := strings.Split(url, "/")
	path := fmt.Sprintf("/repos/%s/%s", info[3], info[4])
	_, err := github.CallGithubApi(path, "GET")
	if err != nil {
		return nil, fmt.Errorf("error in coordinator while calling github api: %s", err)
	}

	var responses []interface{}
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

		server := &mcpserver.AgentMCPServer{
			ServerId:  agentType,
			Transport: transport,
		}

		ac.McpClients[agentType] = client
		ac.McpServers[agentType] = server
		server.RegisterTool(agentType, constants.ServerToToolsMapping[agentType])
		fmt.Println("registered tool on server")
		server.Start(ctx)
		response, err := constants.AgentTypeToFunctionMapping[agentType](client, server, info[3], info[4], callLLM, ctx, responses)

		if err != nil {
			return nil, fmt.Errorf("error occured in the %s agent: %s", agentType, err)
		}
		responses = append(responses, response)
		transport.Close()
		cancel()
	}
	ac.TransportManager.CloseAll()
	return responses, nil
}
