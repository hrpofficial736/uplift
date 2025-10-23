package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/api"
	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/hrpofficial736/uplift/server/internal/utils"
)

func main() {
	config.ConfigLoad()
	formattedPort := ":" + config.Cfg.Port

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := utils.ConnectDatabase(ctx)
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connection established.")

	mux := http.NewServeMux()
	api.RegisterRouter(mux, pool)

	handler := api.MiddleWare(mux)

	log.Printf("🚀 Server running on port %s", formattedPort)
	if err := http.ListenAndServe(formattedPort, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
