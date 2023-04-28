package main

import (
	"fmt"
	"github.com/Tyz3/go-nym"
	"github.com/Tyz3/go-nym/response"
	"github.com/Tyz3/go-nym/tags"
	"os"
	"os/signal"
)

func main() {
	// Kill signals processing
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	// Init the client via server credentials
	client := nym.NewClient("ws://127.0.0.1:1977")

	// Dial a connection to the server
	if err := client.Dial(); err != nil {
		panic(err)
	}
	// Closing the client upon completion of main()
	defer client.Close()

	// Start reading WS messages
	go client.ListenAndServe()

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
	}()

	// Request your own address
	if err := client.SendRequestAsText(nym.NewGetSelfAddress()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send a message
	addr := "9wxVvU4mbk4ZpWLiKEM45Bw3ednUVsfTyooTU3jr6iDR.HcJdxJrPDqWfw6TiauLZEC3mKjFByzFGEtFvDPe2pKW3@Emswx6KXyjRfq1c2k4d4uD2e6nBSbH1biorCZUei8UNS"
	r := nym.NewSend("Mix it up!", addr)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send an anonymous message
	addr = "9wxVvU4mbk4ZpWLiKEM45Bw3ednUVsfTyooTU3jr6iDR.HcJdxJrPDqWfw6TiauLZEC3mKjFByzFGEtFvDPe2pKW3@Emswx6KXyjRfq1c2k4d4uD2e6nBSbH1biorCZUei8UNS"
	replySurbs := 1
	r = nym.NewSendAnonymous("Enjoy your anonymous!", addr, replySurbs)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Reply to an anonymous message
	senderTag := "HJNPCE41NncWGnNSC5w6Cq"
	r = nym.NewReply("Pong.", senderTag)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Waiting for the kill or interrupt signal
	<-interrupt
	fmt.Println("Done.")
}
