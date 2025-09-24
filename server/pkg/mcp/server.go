package mcp

import (
	"context"
	"encoding/json"
)

func (s *AgentMCPServer) RegisterTool (name string, handler func (map[string]interface{}) (interface{}, error)) {
	if s.tools == nil {
		s.tools = make(map[string]func (map[string]interface{}) (interface{}, error))
	}

	s.tools[name] = handler;
}


func (s *AgentMCPServer) Start (ctx context.Context) {
	go func () {
		for {
			select {
			case <-ctx.Done():
				return;
			default:
				reqBytes, _ := s.transport.ReceiveFromClient(ctx)
				var req JSONRPCRequest
				json.Unmarshal(reqBytes, &req);


				if handler, ok := s.tools[req.Method]; ok {
					result, err := handler(req.Params.(map[string]interface{}))
					var response JSONRPCResponse
					response.JSONRPC = "2.0"
					response.ID = req.ID
					
					if err != nil {
						response.Error = &RPCError{
							Code: -32000,
							Message: err.Error(),
						}
					} else {
						b, _ := json.Marshal(result)
						response.Result = b
					}
					responseBytes, _ := json.Marshal(response)
					s.transport.SendToClient(ctx, responseBytes);
				}
			}
		}
	} ()
}