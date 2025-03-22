# About
A tool for reading data from the mixnet contract of nym.com 

# Features

- [x] Query methods
- [ ] Exec methods

## Installation

Download:
```bash
go get github.com/craftdome/go-nym/mixnet@v1.5.1-rc6
```

Import:
```go
import "github.com/craftdome/go-nym/mixnet"
```

# Using

```go
package main

import (
	"context"
	"github.com/craftdome/go-nym/mixnet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	grpcAddr := "nym-grpc.polkachu.com:15390"
	contract := "n17srjznxl9dvzdkpwpw24gg668wc73val88a6m5ajg6ankwvz9wtst0cznr"

	// Init the grpc connection
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// Creating contract client
	client := mixnet.New(conn, contract)

	ctx := context.Background()

	// Getting contract version
	contractVersion, err := client.Contract.GetVersion(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Contract Version: %s", contractVersion.BuildVersion)

	// Getting contract state params
	stateParams, err := client.Contract.GetStateParams(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", stateParams)

	// Getting current interval status of the mixnet
	status, err := client.Contract.GetIntervalStatus(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Epoch started at: %s", status.Interval.CurrentEpochStart.Format(time.DateTime))
}
```

# Support the developer (Nodes)

In the list below you will find developer's nodes. If you are looking for stake your funds, you can take a closer look at my options.

Nymesis Explorer: https://nymesis.vercel.app/?q=3826
