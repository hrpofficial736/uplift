package mcp

import (
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)


func CheckPoint (prompt string, agents []string, callLLM func (string) (types.Response, error)) (types.Response, error) {
	return callLLM(utils.GetCheckPointSystemPrompt(prompt, agents))
}