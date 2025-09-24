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




type LLMResponse struct {
	Content string
	ToolCall string
}
