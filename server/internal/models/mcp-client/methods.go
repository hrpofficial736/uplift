package mcpclient

import (
	"fmt"

	mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"
)

func (c *AgentMCPClient) SetTransport(transport *mcptransport.InMemoryTransport) {
	c.Transport = transport
}

func (c *AgentMCPClient) Initialize() error {
	if c.Transport == nil {
		return fmt.Errorf("transport not set")
	}

	c.Initialized = true
	return nil
}
