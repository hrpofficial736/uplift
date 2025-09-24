package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/hrpofficial736/uplift/server/internal/services/types"
)

func CallLLM (prompt string) (string, error) {
	if (prompt == "") {
		return "", fmt.Errorf("invalid prompt");
	}

	cfg := config.ConfigLoad();
	geminiBaseUrl := cfg.GeminiBaseUrl;
	geminiModel := cfg.GeminiModel;
	geminiApiKey := cfg.GeminiAPIKey;


	requestUrl := fmt.Sprintf("%s/%s:generateContent?key=%s", geminiBaseUrl, geminiModel, geminiApiKey);

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

	jsonBody, err := json.Marshal(body);

	if err != nil {
		log.Fatal()
	}

	response, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonBody));


	if err != nil {
		return "", nil;
	}

	data, err := io.ReadAll(response.Body);

	if err != nil {
		return "", err;
	}

	var result map[string]interface{};

	if err := json.Unmarshal(data, &result); err != nil {
		return "", err
	}

	candidates := result["candidates"].([]interface{});

	if len(candidates) == 0 {
		return "", fmt.Errorf("no candidates returned")
	}

	content := candidates[0].(map[string]interface{})["content"].(map[string]interface{});

	parts := content["parts"].([]interface{});

	text := parts[0].(map[string]interface{})["text"].(string);

	return text, nil;

}