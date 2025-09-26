package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func NewAgentMCPClient (agentType string, callLLM func (string) (utils.Response, error)) *AgentMCPClient {
	ctx, cancel := context.WithCancel(context.Background());

	return &AgentMCPClient{
		agentType: agentType,
		callLLM: callLLM,
		ctx: ctx,
		cancel: cancel,
	}
}


func (c *AgentMCPClient) SetTransport (transport *InMemoryTransport) {
	c.transport = transport;
}

func (c *AgentMCPClient) Initialize () error {
	if c.transport == nil {
		return fmt.Errorf("transport not set");
	}

	c.initialized = true
	return nil	
}


func (c *AgentMCPClient) CallTool (ctx context.Context, method string, params map[string]interface{}) (string, error) {
	if !c.initialized {
		return "", errors.New("client not initialized")
	}

	c.requestID++

	req := JSONRPCRequest{
		JSONRPC: "2.0",
		Method: method,
		Params: params,
		ID: c.requestID,
	}

	reqBytes, _ := json.Marshal(req)
	c.transport.Send(ctx, reqBytes)

	responseBytes, _ := c.transport.Receive(ctx)

	var response JSONRPCResponse
	json.Unmarshal(responseBytes, &response)

	if response.Error != nil {
		return "", errors.New(response.Error.Message)
	}

	var result string
	json.Unmarshal(response.Result, &result)

	return result, nil
}


