package utils

import "log"

type Response struct {
	Text      string
	ToolCalls interface{}
}

func FormatModelResponse(result map[string]interface{}) Response {
	candidates, ok := result["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		log.Fatal("No candidates returned!")
	}

	content, ok := candidates[0].(map[string]interface{})["content"].(map[string]interface{})
	if !ok {
		log.Fatal("Invalid candidate content structure")
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		log.Fatal("No parts found in candidate content")
	}

	part, ok := parts[0].(map[string]interface{})
	if !ok {
		log.Fatal("Invalid part structure")
	}

	text, _ := part["text"].(string)

	var toolCalls interface{}
	if fc, exists := part["function_calls"]; exists {
		toolCalls = fc
	}

	return Response{
		Text:      text,
		ToolCalls: toolCalls,
	}
}
