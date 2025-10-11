package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	GithubAccessToken string
	GithubBaseUrl     string
	GeminiBaseUrl     string
	GeminiModel       string
	GeminiAPIKey      string
	ClientUrl         string
}

func ConfigLoad() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		Port:              getEnv("PORT", "7777"),
		GithubAccessToken: getEnv("GITHUB_ACCESS_TOKEN", "token"),
		GithubBaseUrl:     getEnv("GITHUB_API_BASE_URL", "github_url"),
		GeminiBaseUrl:     getEnv("GEMINI_BASE_URL", "gemini_url"),
		GeminiModel:       getEnv("GEMINI_MODEL", "gemini-2.0-flash"),
		GeminiAPIKey:      getEnv("GEMINI_API_KEY", "gemini_api_key"),
		ClientUrl:         getEnv("CLIENT_URL", "client_url"),
	}
	log.Println("Config loaded!")
	return cfg
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
