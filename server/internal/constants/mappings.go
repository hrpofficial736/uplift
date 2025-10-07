package constants

import (
	"context"

	"github.com/hrpofficial736/uplift/server/internal/agents"
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	mcpserver "github.com/hrpofficial736/uplift/server/internal/models/mcp-server"
	"github.com/hrpofficial736/uplift/server/internal/services/github"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

var AgentTypeToFunctionMapping = map[string]func(*mcpclient.AgentMCPClient, *mcpserver.AgentMCPServer, string, string, func(string) (types.Response, error), context.Context, []interface{}) (interface{}, error){
	"security":    agents.SecurityCritic,
	"performance": agents.MaintainabilityCritic,
	"quality":     agents.QualityCritic,
	"mentor":      agents.Mentor,
}

var ServerToToolsMapping = map[string]func(string, string) ([]interface{}, error){
	"security":        github.CheckRepoSecurity,
	"maintainability": github.CheckRepoMaintainability,
	"quality":         github.CheckForQuality,
}
