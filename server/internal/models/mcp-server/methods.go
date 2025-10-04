package mcpserver

import (
	"context"
	"encoding/json"

	jsonrpctypes "github.com/hrpofficial736/uplift/server/internal/models/json-rpc-types"
)


func (s *AgentMCPServer) RegisterTool (name string, handler func (map[string]interface{}) (interface{}, error)) {
	if s.Tools == nil {
		s.Tools = make(map[string]func (map[string]interface{}) (interface{}, error))
	}

	s.Tools[name] = handler;
}


func (s *AgentMCPServer) Start (ctx context.Context) {
	go func () {
		for {
			select {
			case <-ctx.Done():
				return;
			default:
				reqBytes, _ := s.Transport.ReceiveFromClient(ctx)
				var req jsonrpctypes.JSONRPCRequest
				json.Unmarshal(reqBytes, &req);


				if handler, ok := s.Tools[req.Method]; ok {
					result, err := handler(req.Params.(map[string]interface{}))
					var response jsonrpctypes.JSONRPCResponse
					response.JSONRPC = "2.0"
					response.ID = req.ID
					
					if err != nil {
						response.Error = &jsonrpctypes.RPCError{
							Code: -32000,
							Message: err.Error(),
						}
					} else {
						b, _ := json.Marshal(result)
						response.Result = b
					}
					responseBytes, _ := json.Marshal(response)
					s.Transport.SendToClient(ctx, responseBytes);
				}
			}
		}
	} ()
}