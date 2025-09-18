package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseUrl string
	GithubAccessToken string
	GithubBaseUrl string
}


func ConfigLoad () *Config {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }
	
	cfg := &Config{
		Port: getEnv("PORT", "7777"),
		DatabaseUrl: getEnv("DATABASE_URL", "hello"),
		GithubAccessToken: getEnv("GITHUB_ACCESS_TOKEN", "token"),
		GithubBaseUrl: getEnv("GITHUB_API_BASE_URL", "github_url"),
	}
	log.Printf("Config loaded, PORT = %s and DATABASE_URL = %s", cfg.Port, cfg.DatabaseUrl);
	return cfg;
}


func getEnv (key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value;
	}

	return defaultValue;
}