package mcp

import (
	"github.com/hrpofficial736/uplift/server/internal/utils"
)


func CheckPoint (prompt string, agents []string, callLLM func (string) (utils.Response, error)) (utils.Response, error) {
	return callLLM(utils.GetCheckPointSystemPrompt(prompt, agents))
}