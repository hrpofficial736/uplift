package mcpserver

import mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"

type AgentMCPServer struct {
	ServerId  string
	Transport *mcptransport.InMemoryTransport
	Tools     map[string]func(string, string) ([]interface{}, error)
}
