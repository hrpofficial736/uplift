package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func CallLLM(prompt string) (utils.Response, error) {
	if prompt == "" {
		return utils.Response{}, fmt.Errorf("invalid prompt")
	}

	cfg := config.ConfigLoad()
	geminiBaseUrl := cfg.GeminiBaseUrl
	geminiModel := cfg.GeminiModel
	geminiApiKey := cfg.GeminiAPIKey

	requestUrl := fmt.Sprintf("%s/%s:generateContent?key=%s", geminiBaseUrl, geminiModel, geminiApiKey)

	body := types.LLMRequestBody{
		Contents: []types.Content{
			{
				Parts: []types.Part{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return utils.Response{}, fmt.Errorf("failed to marshal body: %w", err)
	}

	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return utils.Response{}, fmt.Errorf("http post failed: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return utils.Response{}, fmt.Errorf("read body failed: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return utils.Response{}, fmt.Errorf("unmarshal failed: %w", err)
	}

	return utils.FormatModelResponse(result), nil
}
