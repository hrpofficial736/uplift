package agents

import (
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

func PerformanceCritic(client *mcpclient.AgentMCPClient, server *mcpserver.AgentMCPServer, url string, callLLM func(string) (types.Response, error)) string {
	return ""
}
