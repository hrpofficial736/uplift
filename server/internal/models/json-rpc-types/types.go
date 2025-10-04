package jsonrpctypes

import "encoding/json"


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