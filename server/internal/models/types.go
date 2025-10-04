package models

type CheckPointResponse struct {
	Valid   bool     `json:"valid"`
	Agents  []string `json:"agents"`
	Message string   `json:"message"`
	Url     string   `json:"url"`
}
