package main

import (
	"context"
	"github.com/craftdome/go-nym/clients/nymapi"
	"golang.org/x/time/rate"
	"log"
	"net/http"
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
