package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	jsonrpctypes "github.com/hrpofficial736/uplift/server/internal/models/json-rpc-types"
)

func (s *AgentMCPServer) RegisterTool(name string, handler func(string, string) ([]interface{}, error)) {
	if s.Tools == nil {
		s.Tools = make(map[string]func(string, string) ([]interface{}, error))
	}

	s.Tools[name] = handler
}

func (s *AgentMCPServer) Start(ctx context.Context) {
	fmt.Println("in mcp server now...")
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				reqBytes, _ := s.Transport.ReceiveFromClient(ctx)
				reqJSON, err := json.Marshal(reqBytes)
				if err != nil {
					log.Fatalf("error in mcp server while converting the incoming request to json: %s", err)
				}
				var req jsonrpctypes.JSONRPCRequest
				if err := json.Unmarshal(reqJSON, &req); err != nil {
					log.Fatalf("error in mcp server while unmarshalling the json to the json rpc request: %s", err)
				}
				fmt.Printf("request recived from the client: %v\n", req)
				if handler, ok := s.Tools[req.Method]; ok {
					result, err := handler(req.Params["owner"], req.Params["repo"])
					if err != nil {
						log.Fatalf("error while tool calling in mcp server: %s", err)
					}
					fmt.Println(result)
					var response jsonrpctypes.JSONRPCResponse
					response.JSONRPC = "2.0"
					response.ID = req.ID

					if err != nil {
						response.Error = &jsonrpctypes.RPCError{
							Code:    -32000,
							Message: err.Error(),
						}
					} else {
						b, _ := json.Marshal(result)
						response.Result = b
					}
					responseBytes, _ := json.Marshal(response.Result)
					s.Transport.SendToClient(ctx, responseBytes)
				}
			}
		}
	}()
}
