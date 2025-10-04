package mcp

import (
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpcoordinator "github.com/hrpofficial736/uplift/server/internal/models/mcp-coordinator"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
)


func NewAgentCoordinator () *mcpcoordinator.AgentCoordinator {
	return &mcpcoordinator.AgentCoordinator{
		TransportManager: NewTransportManager(),
		McpClients: make(map[string]*mcpclient.AgentMCPClient),
		McpServers: make(map[string]*mcpserver.AgentMCPServer),
	}
}

