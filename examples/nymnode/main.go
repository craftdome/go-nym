package main

import (
	"context"
	"github.com/craftdome/go-nym/nymnode"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	client := &http.Client{}
	limiter := rate.NewLimiter(rate.Inf, 0) // requests limiter

	host := "nym-exit.hkg2.craftdome.app:8080"

	// Node client
	node, err := nymnode.New(
		ctx, host,
		nymnode.WithRateLimiter(limiter),
		nymnode.WithCustomClient(client),
	)
	if err != nil {
		log.Fatal(err)
	}

	desc, err := node.GetDescription(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Description: %+v", desc)
	// Description: {Moniker:3826.4 project Website:https://github.com/craftdome SecurityContact:@craftdome Details:Software Developer | NYM Dev Grantee}
}
