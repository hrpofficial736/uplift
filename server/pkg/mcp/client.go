package mcp

import (
	"context"
	mcpclient "github.com/hrpofficial736/uplift/server/internal/models/mcp-client"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

func NewAgentMCPClient (agentType string, callLLM func (string) (types.Response, error)) *mcpclient.AgentMCPClient {
	ctx, cancel := context.WithCancel(context.Background());

	return &mcpclient.AgentMCPClient{
		AgentType: agentType,
		CallLLM: callLLM,
		Ctx: ctx,
		Cancel: cancel,
	}
}





