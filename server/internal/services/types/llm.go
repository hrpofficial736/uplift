package types

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}
type LLMRequestBody struct {
	Contents []Content `json:"contents"`
}

type Response struct {
	Text      string      `json:"text"`
	ToolCalls interface{} `json:"toolCalls"`
}

type AgentResponse struct {
	Data  Response `json:"data"`
	Agent string   `json:"agent,omitempty"`
}
