package mcp

import "encoding/json"

func mustMarshal (v interface{}) json.RawMessage {
	b, _ := json.Marshal(v);
	return json.RawMessage(b);
}