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

func (ac *AgentCoordinator) AddAgentAndGetAgentResponse(agents []string, callLLM func(string) (types.Response, error), prompt string, url string) ([]interface{}, interface{}, error) {
	fmt.Println("inside the coordinator with")
	info := strings.Split(url, "/")
	fmt.Println(info)
	path := fmt.Sprintf("/repos/%s/%s", info[3], info[4])
	fmt.Println(path)
	_, err := github.CallGithubApi(path, "GET")
	if err != nil {
		return nil, nil, fmt.Errorf("error in coordinator while calling github api: %s", err)
	}

	var responses []interface{}
	fmt.Println(len(agents))
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
		fmt.Println("assigned client and server, now about to register tools on the server!")
		server.RegisterTool(agentType, constants.ServerToToolsMapping[agentType])
		fmt.Println("registered tool on server")
		server.Start(ctx)
		response, err := constants.AgentTypeToFunctionMapping[agentType](client, server, info[3], info[4], callLLM, ctx, responses)

		if err != nil {
			return nil, nil, fmt.Errorf("error occured in the %s agent: %s", agentType, err)
		}
		responses = append(responses, response)
		transport.Close()
		cancel()
	}
	ac.TransportManager.CloseAll()
	return responses, map[string]string{
		"ownerName": info[3],
		"repoName":  info[4],
	}, nil
}
