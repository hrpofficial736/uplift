package mcpcoordinator

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/hrpofficial736/uplift/server/internal/constants"
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/github"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

func (ac *AgentCoordinator) AddAgentAndGetAgentResponse(agents []string, callLLM func(string) (types.Response, error), prompt string, url string) ([]interface{}, interface{}, error) {
	fmt.Println("inside the coordinator")
	info := strings.Split(url, "/")
	fmt.Println(info)
	path := fmt.Sprintf("/repos/%s/%s", info[3], info[4])
	fmt.Println(path)
	_, err := github.CallGithubApi(path, "GET")
	if err != nil {
		return nil, nil, fmt.Errorf("error in coordinator while calling github api: %s", err)
	}

	var (
		wg               sync.WaitGroup
		mu               sync.Mutex
		responses        []interface{}
		responsesChannel = make(chan interface{}, len(agents))
		errorChannel     = make(chan error, len(agents))
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, agentType := range agents {
		wg.Add(1)
		go func(agentType string) {
			defer wg.Done()

			agentCtx, agentCancel := context.WithCancel(ctx)
			defer agentCancel()
			transport := ac.TransportManager.CreateTransport(agentType)

			client := &mcpclient.AgentMCPClient{
				AgentType: agentType,
				CallLLM:   callLLM,
				Transport: transport,
				Ctx:       agentCtx,
				Cancel:    agentCancel,
			}

			server := &mcpserver.AgentMCPServer{
				ServerId:  agentType,
				Transport: transport,
			}
			mu.Lock()
			ac.McpClients[agentType] = client
			ac.McpServers[agentType] = server
			mu.Unlock()
			fmt.Println("assigned client and server, now about to register tools on the server!")
			server.RegisterTool(agentType, constants.ServerToToolsMapping[agentType])
			fmt.Println("registered tool on server")
			server.Start(agentCtx)
			response, err := constants.AgentTypeToFunctionMapping[agentType](client, server, info[3], info[4], callLLM, agentCtx, responses)

			if err != nil {
				errorChannel <- fmt.Errorf("error while calling agents: %s", err)
				cancel()
				return
			}
			responsesChannel <- response
			transport.Close()
		}(agentType)
	}

	wg.Wait()
	close(responsesChannel)
	close(errorChannel)

	for response := range responsesChannel {
		responses = append(responses, response)
	}

	var combinedErr error
	for err := range errorChannel {
		combinedErr = errors.Join(combinedErr, err)
	}
	ac.TransportManager.CloseAll()

	if combinedErr != nil {
		return nil, nil, combinedErr
	}
	return responses, map[string]string{
		"ownerName": info[3],
		"repoName":  info[4],
	}, nil
}
