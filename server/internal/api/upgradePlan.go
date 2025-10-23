package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/hrpofficial736/uplift/server/internal/services/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/webhook"
)

func upgradePlan(pool *pgxpool.Pool) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("in upgrade plan handler")
		const maxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(res, req.Body, maxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("error reading request body")
			res.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		stripeWebhookSecret := config.Cfg.StripeWebhookSecret

		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"), stripeWebhookSecret)

		if err != nil {
			fmt.Printf("error constructing webhook event: %s\n", err)
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if event.Type == "checkout.session.completed" {
			var session stripe.CheckoutSession
			fmt.Println("checkout session completed")
			if err := json.Unmarshal(event.Data.Raw, &session); err == nil {
				userId := session.ClientReferenceID
				log.Printf("payment success for user: %s", userId)
				fmt.Println(userId)
				rows, err := database.QueryDatabase(context.Background(), pool,
					`UPDATE "Users" SET plan = $1, plan_upgraded_at = now() WHERE id = $2`, "PRO", userId)
				if err != nil {
					fmt.Println("database query error")
					http.Error(res, fmt.Sprintf("database query error: %v", err), http.StatusInternalServerError)
					return
				}
				defer rows.Close()

				res.WriteHeader(http.StatusOK)
			}
		}
	}
}
