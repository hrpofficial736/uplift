package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	GithubAccessToken   string
	GithubBaseUrl       string
	GeminiBaseUrl       string
	GeminiModel         string
	GeminiAPIKey        string
	ClientUrl           string
	DatabaseUrl         string
	SupabaseJWTSecret   string
	StripeSecretKey     string
	StripeWebhookSecret string
}

var Cfg *Config

func ConfigLoad() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	Cfg = &Config{
		Port:                getEnv("PORT", "7777"),
		GithubAccessToken:   getEnv("GITHUB_ACCESS_TOKEN", "token"),
		GithubBaseUrl:       getEnv("GITHUB_API_BASE_URL", "github_url"),
		GeminiBaseUrl:       getEnv("GEMINI_BASE_URL", "gemini_url"),
		GeminiModel:         getEnv("GEMINI_MODEL", "gemini-2.0-flash"),
		GeminiAPIKey:        getEnv("GEMINI_API_KEY", "gemini_api_key"),
		ClientUrl:           getEnv("CLIENT_URL", "client_url"),
		DatabaseUrl:         getEnv("DATABASE_URL", "database_url"),
		SupabaseJWTSecret:   getEnv("SUPABASE_JWT_SECRET", "jwt_secret"),
		StripeSecretKey:     getEnv("STRIPE_SECRET_KEY", "stripe_secret_key"),
		StripeWebhookSecret: getEnv("STRIPE_WEBHOOK_SECRET", "stripe_webhook_secret"),
	}
	log.Println("Config loaded!")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
