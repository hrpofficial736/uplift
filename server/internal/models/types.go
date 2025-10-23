package models

type CheckPointResponse struct {
	Valid   bool     `json:"valid"`
	Agents  []string `json:"agents"`
	Message string   `json:"message"`
	Url     string   `json:"url"`
}

type Request struct {
	Email  string `json:"email"`
	Prompt string `json:"prompt"`
}

type Response struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	RepoInfo interface{} `json:"repoInfo"`
	Reviewed bool        `json:"reviewed"`
}

type User struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Plan             string `json:"plan"`
	Prompts          int    `json:"prompts"`
	Plan_upgraded_at string `json:"plan_upgraded_at"`
	CreatedAt        string `json:"created_at"`
}
