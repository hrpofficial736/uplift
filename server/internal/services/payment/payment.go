package payment

import (
	"fmt"
	"log"

	"github.com/hrpofficial736/uplift/server/internal/config"
	"github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/checkout/session"
)

func HandleCreateCheckoutSession(userId string) (interface{}, error) {
	cfg := config.Cfg
	stripe.Key = cfg.StripeSecretKey
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("inr"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Pro Plan"),
					},
					UnitAmount: stripe.Int64(49900),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL:        stripe.String(cfg.ClientUrl),
		CancelURL:         stripe.String(cfg.ClientUrl),
		ClientReferenceID: stripe.String(userId),
		Metadata: map[string]string{
			"userId": userId,
		},
	}
	s, err := session.New(params)
	if err != nil {
		log.Println("Stripe session error: ", err)
		return nil, fmt.Errorf("stripe session error: %s", err)
	}

	return s.URL, nil
}
