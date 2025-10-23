package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
)

func CallGithubApi(path string, method string) (interface{}, error) {
	cfg := config.Cfg
	githubBaseUrl := cfg.GithubBaseUrl
	token := cfg.GithubAccessToken
	requestUrl := fmt.Sprintf("%s%s", githubBaseUrl, path)
	request, err := http.NewRequest(method, requestUrl, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating the request: %s", err)
	}

	request.Header.Set("Authorization", "token "+token)
	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("User-Agent", "Uplift")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error while calling the github api: %s", err)
	}

	var fResponse interface{}
	if err := json.NewDecoder(response.Body).Decode(&fResponse); err != nil {
		return nil, fmt.Errorf("error while converting json response into interface: %s", err)
	}
	return fResponse, nil
}
