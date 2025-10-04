package utils

import (
	"github.com/hrpofficial736/uplift/server/internal/agents"
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

var AgentTypeToFunctionMapping = map[string]func(*mcpclient.AgentMCPClient, *mcpserver.AgentMCPServer, string, func(string) (types.Response, error)) string{
	"security":    agents.SecurityCritic,
	"performance": agents.PerformanceCritic,
	"quality":     agents.Perfectionist,
	"mentor":      agents.Mentor,
}
