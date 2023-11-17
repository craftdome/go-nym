package main

import (
	"fmt"
	"github.com/craftdome/go-nym"
	"github.com/craftdome/go-nym/response"
	"github.com/craftdome/go-nym/tags"
	"os"
	"os/signal"
)

var (
	done = make(chan struct{}, 1)
)

func main() {
	// Kill signals processing
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	// Init the client via server credentials
	client := nym.NewClient("ws://192.168.88.4:1977")

	// Dial a connection to the server
	if err := client.Dial(); err != nil {
		panic(err)
	}

	// Start reading WS messages
	go func() {
		if err := client.ListenAndServe(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	go func() {
		// Incoming Message Channel
		for message := range client.Messages() {
			switch message.Type() {
			case tags.Error:
				msg := message.(*response.Error)
				fmt.Printf("Error: %s\n", msg.Message)
			case tags.SelfAddress:
				msg := message.(*response.SelfAddress)
				fmt.Printf("SelfAddress: %s\n", msg.Address)
			case tags.Received:
				msg := message.(*response.Received)
				fmt.Printf("Received: %s, SenderTag: %s\n", msg.Message, msg.SenderTag)
			}
		}

		fmt.Println("Closed")
		done <- struct{}{}
	}()

	// Request your own address
	if err := client.SendRequestAsText(nym.NewGetSelfAddress()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send a message
	addr := "2w2mvQzGHuzXdz1pQSvTWXiqZe26Z2BKNkFTQ5g7MuLi.DfkhfLipgtuRLAWWHx74iGkJWCpM6U5RFwaJ3FUaMicu@HWdr8jgcr32cVGbjisjmwnVF4xrUBRGvbw86F9e3rFzS"
	r := nym.NewSend("Mix it up!", addr)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send an anonymous message
	addr = "2w2mvQzGHuzXdz1pQSvTWXiqZe26Z2BKNkFTQ5g7MuLi.DfkhfLipgtuRLAWWHx74iGkJWCpM6U5RFwaJ3FUaMicu@HWdr8jgcr32cVGbjisjmwnVF4xrUBRGvbw86F9e3rFzS"
	replySurbs := 1
	r = nym.NewSendAnonymous("Enjoy your anonymous!", addr, replySurbs)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Reply to an anonymous message
	senderTag := "7vv2LmF9M6EwQRrmCiCJhr"
	r = nym.NewReply("Pong.", senderTag)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Waiting for the kill or interrupt signal
	<-interrupt
	// Closing the client
	if err := client.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// Waiting for the done signal
	<-done
	fmt.Println("Done.")
}
