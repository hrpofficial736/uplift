package main

import (
	"log"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/hrpofficial736/uplift/server/internal/api"
)

func main () {
	cfg := config.ConfigLoad();
	formattedPort := ":" + cfg.Port;
	
	mux := http.NewServeMux();
	
	api.RegisterRouter(mux);

	log.Printf("server is listening at %s...\n", cfg.Port);
	log.Fatal(http.ListenAndServe(formattedPort, mux));


}