package models

type CheckPointResponse struct {
	Valid   bool     `json:"valid"`
	Agents  []string `json:"agents"`
	Message string   `json:"message"`
	Url     string   `json:"url"`
}

type Request struct {
	Prompt string   `json:"prompt"`
	Agents []string `json:"agents"`
}

type Response struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	RepoInfo interface{} `json:"repoInfo"`
	Reviewed bool        `json:"reviewed"`
}
