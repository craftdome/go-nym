package main

import (
	"context"
	"github.com/craftdome/go-nym/contracts/mixnet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	grpcAddr := "nym-grpc.polkachu.com:15390"
	contract := "n17srjznxl9dvzdkpwpw24gg668wc73val88a6m5ajg6ankwvz9wtst0cznr"

	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := mixnet.NewQueryClient(conn, contract)

	ctx := context.Background()

	stateParams, err := client.Contract.GetStateParams(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", stateParams)

	status, err := client.Contract.GetIntervalStatus(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Epoch started at: %s", status.Interval.CurrentEpochStart.Format(time.DateTime))
}
