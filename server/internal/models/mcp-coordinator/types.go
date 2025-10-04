package mcpcoordinator

import (
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"
)

type AgentCoordinator struct {
	TransportManager *mcptransport.TransportManager
	McpClients map[string]*mcpclient.AgentMCPClient
	McpServers map[string]*mcpserver.AgentMCPServer
}