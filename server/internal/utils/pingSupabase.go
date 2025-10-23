package utils

import (
	"fmt"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
)

func PingSupabase() {
	cfg := config.Cfg
	req, err := http.NewRequest("GET", cfg.SupabaseUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("apikey", cfg.SupabaseServiceKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error pinging Supabase:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Pinged Supabase:", resp.Status)
}
