package mcpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	
	jsonrpctypes "github.com/hrpofficial736/uplift/server/internal/models/json-rpc-types"
	mcptransport "github.com/hrpofficial736/uplift/server/internal/models/mcp-transport"
)

func (c *AgentMCPClient) SetTransport (transport *mcptransport.InMemoryTransport) {
	c.Transport = transport;
}

func (c *AgentMCPClient) Initialize () error {
	if c.Transport == nil {
		return fmt.Errorf("transport not set");
	}

	c.Initialized = true
	return nil	
}


func (c *AgentMCPClient) CallTool (ctx context.Context, method string, params map[string]interface{}) (string, error) {
	if !c.Initialized {
		return "", errors.New("client not initialized")
	}

	c.RequestID++

	req := jsonrpctypes.JSONRPCRequest{
		JSONRPC: "2.0",
		Method: method,
		Params: params,
		ID: c.RequestID,
	}

	reqBytes, _ := json.Marshal(req)
	c.Transport.Send(ctx, reqBytes)

	responseBytes, _ := c.Transport.Receive(ctx)

	var response jsonrpctypes.JSONRPCResponse
	json.Unmarshal(responseBytes, &response)

	if response.Error != nil {
		return "", errors.New(response.Error.Message)
	}

	var result string
	json.Unmarshal(response.Result, &result)

	return result, nil
}