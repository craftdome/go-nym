# About
A client implementation of nym-node API

# Features

- [x] Query methods
- [ ] Exec methods

## Installation

Download:
```bash
go get github.com/craftdome/go-nym/nymnode@v1.7.0-rc5
```

Import:
```go
import "github.com/craftdome/go-nym/nymnode"
```

# Using

```go
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

```

# Support the developer (Nodes)

In the list below you will find developer's nodes. If you are looking for stake your funds, you can take a closer look at my options.

Nymesis Explorer: https://nymesis.vercel.app/?q=3826
