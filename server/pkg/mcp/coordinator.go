package mcp

import (
	"context"
	"fmt"

	"github.com/hrpofficial736/uplift/server/internal/utils"
)


func NewAgentCoordinator () *AgentCoordinator {
	return &AgentCoordinator{
		transportManager: NewTransportManager(),
		mcpClients: make(map[string]*AgentMCPClient),
		mcpServers: make(map[string]*AgentMCPServer),
	}
}


func (ac *AgentCoordinator) AddAgent (agents []string, callLLM func (string) (utils.Response, error), prompt string) (interface{}, error) {
		
	for _, agentType := range agents {
		transport := ac.transportManager.CreateTransport(agentType)
		ctx, cancel := context.WithCancel(context.Background())

		client := &AgentMCPClient{
			agentType: agentType,
			callLLM: callLLM,
			transport: transport,
			ctx: ctx,
			cancel: cancel,
		}

		client.Initialize()

		server := &AgentMCPServer{
			serverId: agentType,
			transport: transport,
		}


		ac.mcpClients[agentType] = client
		ac.mcpServers[agentType] = server
	}

	return nil, fmt.Errorf("error aa rha h")
} 