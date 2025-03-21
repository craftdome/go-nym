# About
This library is designed to simplify interaction with the NYM protocol for [nym-client](https://nimtech.net/docs/clients/overview.html#the-websocket-client). It implements a basic set of commands for sending and receiving messages to/from mixnet.

# Features

- [x] Using [Gorilla WebSocket](https://github.com/gorilla/websocket )
- [x] Text Protocol (JSON) support
- [ ] Binary Protocol support
- [ ] Support for the user protocol in the body of the binary message

# Preparation for use

1. The library requires an active connection to the websocket client. The simple installation and launch of the `nym-client` is described in the official docs -> [link](https://nymtech.net/docs/clients/websocket/setup.html)

2. Tested with nym-client version `1.1.32`

3. Importing a dependency

Download:
```
go get github.com/craftdome/go-nym/clients/wsc@v1.0.2
```

Import:
```go
import "github.com/craftdome/go-nym/clients/wsc"
```

# Using

### Tips

- Remember, if you plan to give access to `nym-client` connect from the outside (for this you should specify the ip of the external network interface of your machine), the `nym-client` doesn't have a connection auth function. 
- Only 1 connection is allowed at a time.
- If you need an external connection, use a local networks ([10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12](https://en.wikipedia.org/wiki/Private_network)) instead of a global unicast one to increase security.

### Example

```go
package main

import (
	"fmt"
	"github.com/craftdome/go-nym/clients/wsc"
	"github.com/craftdome/go-nym/clients/wsc/response"
	"github.com/craftdome/go-nym/clients/wsc/tags"
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
	client := wsc.New("ws://192.168.88.4:1977")

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
	if err := client.SendRequestAsText(wsc.NewGetSelfAddress()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send a message
	addr := "2w2mvQzGHuzXdz1pQSvTWXiqZe26Z2BKNkFTQ5g7MuLi.DfkhfLipgtuRLAWWHx74iGkJWCpM6U5RFwaJ3FUaMicu@HWdr8jgcr32cVGbjisjmwnVF4xrUBRGvbw86F9e3rFzS"
	r := wsc.NewSend("Mix it up!", addr)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Send an anonymous message
	addr = "2w2mvQzGHuzXdz1pQSvTWXiqZe26Z2BKNkFTQ5g7MuLi.DfkhfLipgtuRLAWWHx74iGkJWCpM6U5RFwaJ3FUaMicu@HWdr8jgcr32cVGbjisjmwnVF4xrUBRGvbw86F9e3rFzS"
	replySurbs := 1
	r = wsc.NewSendAnonymous("Enjoy your anonymous!", addr, replySurbs)
	if err := client.SendRequestAsText(r); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Reply to an anonymous message
	senderTag := "7vv2LmF9M6EwQRrmCiCJhr"
	r = wsc.NewReply("Pong.", senderTag)
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
```

# Support the developer (Nodes)

In the list below you will find developer's nodes. If you are looking for stake your funds, you can take a closer look at my options.

Nymesis Explorer: https://nymesis.vercel.app/?q=3826
