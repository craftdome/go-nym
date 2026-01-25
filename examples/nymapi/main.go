package main

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/time/rate"

	"github.com/craftdome/go-nym/nymapi"
)

func main() {
	ctx := context.Background()
	client := &http.Client{}
	limiter := rate.NewLimiter(rate.Inf, 0) // requests limiter

	host := "validator.nymtech.net"

	api, err := nymapi.New(
		ctx,
		host,
		nymapi.WithCustomClient(client),
		nymapi.WithRateLimiter(limiter),
	)
	if err != nil {
		log.Fatal(err)
	}

	health, err := api.Health(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("health status: %+v", health)
}
