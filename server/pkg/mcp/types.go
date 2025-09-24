package mcp

import (
	"context"
	"encoding/json"
	"sync"
)


type InMemoryTransport struct {
	serverId string
	clientChannel chan []byte
	serverChannel chan []byte
	closed bool
	mu sync.RWMutex
}


type TransportManager struct {
	transports map[string]*InMemoryTransport
	mu sync.RWMutex
}

type JSONRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params,omitempty"`
	ID int64 `json:"id"`
}


type RPCError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}


type JSONRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result json.RawMessage `json:"result,omitempty"`
	Error *RPCError `json:"error,omitempty"`
	ID int64 `json:"id"`
}

type AgentMCPClient struct {
	agentType string
	callLLM func (string) (string, error)
	transport *InMemoryTransport
	initialized bool
	requestID int64
	ctx context.Context
	cancel context.CancelFunc
}

type AgentMCPServer struct {
	serverId string
	transport *InMemoryTransport
	tools map[string]func(map[string]interface{}) (interface{}, error)
} 


type AgentCoordinator struct {
	transportManager *TransportManager
	mcpClients map[string]*AgentMCPClient
	mcpServers map[string]*AgentMCPServer
}