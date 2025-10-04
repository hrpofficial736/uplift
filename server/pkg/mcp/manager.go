package mcp

import (

	mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"
)


func NewTransportManager () *mcptransport.TransportManager {
	return &mcptransport.TransportManager{
		Transports: make(map[string]*mcptransport.InMemoryTransport),
	}
}


