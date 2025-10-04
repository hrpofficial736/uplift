package mcpclient

import (
	"context"

	mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

type AgentMCPClient struct {
	AgentType string
	CallLLM func (string) (types.Response, error)
	Transport *mcptransport.InMemoryTransport
	Initialized bool
	RequestID int64
	Ctx context.Context
	Cancel context.CancelFunc
}