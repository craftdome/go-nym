# About
A client implementation of nym-node API

# Features

- [x] Query methods
- [ ] Exec methods

## Installation

Download:
```bash
go get github.com/craftdome/go-nym/clients/nymapi@v1.1.54-rc1
```

Import:
```go
import "github.com/craftdome/go-nym/clients/nymapi"
```

# Using

```go
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
```

# Support the developer (Nodes)

In the list below you will find developer's nodes. If you are looking for stake your funds, you can take a closer look at my options.

Nymesis Explorer: https://nymesis.vercel.app/?q=3826
