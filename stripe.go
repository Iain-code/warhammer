package main

import (
    "log"
    "net/http"
	"os"
	
    "github.com/stripe/stripe-go/v82"
    "github.com/stripe/stripe-go/v82/checkout/session"
)

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {

  stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
  domain := "http://localhost:8080"
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        // Provide the exact Price ID (for example, price_1234) of the product you want to sell
        Price: stripe.String("{{PRICE_ID}}"),
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    SuccessURL: stripe.String(domain + "?success=true"),
    CancelURL: stripe.String(domain + "?canceled=true"),
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
  }

  http.Redirect(w, r, s.URL, http.StatusSeeOther)
}