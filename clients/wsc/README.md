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
go get github.com/craftdome/go-nym
```

Import:
```go
import "github.com/craftdome/go-nym/wsc"
```

# Using

### Tips

- Remember, if you plan to give access to `nym-client` connect from the outside (for this you should specify the ip of the external network interface of your machine), the `nym-client` doesn't have a connection auth function. 
- Only 1 connection is allowed at a time.
- If you need an external connection, use a local networks ([10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12](https://en.wikipedia.org/wiki/Private_network)) instead of a global unicast one to increase security.

### Initialization

1. First, you need to `init` the connection client with the nym-client address. We copy the connection `address:port` from the console after running the nym-client. By default it's `localhost:1977`, or in my case `192.168.88.4:1977`.

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L21-L22

2. Connection establishing with the `nym-client`.

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L24-L27

### Reading messages

3. We turn on listening to incoming messages, which we then extract through the `Messages()` chan.

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L29-L54

### Sending messages

4. Getting your nym-client address (SelfAddress).

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L56-L59

5. Sending a message (Send).

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L61-L66

6. Sending a SURB message to receive an anonymous response (SendAnonymous).

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L68-L74

7. Sending a reply to the SendAnonymous message (Reply).

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L76-L81

### Closing the connection correctly

8. We close the connection with the `nym-client` after the interrupt (sigint/sigkill) and wait for the reading gorutine sent `done` signal.

https://github.com/craftdome/go-nym/blob/7db430e3a036044c42c447c79d8fbeae7a068806/examples/clients/wsc/main.go#L83-L91

# Support the developer (Nodes)

In the list below you will find developer's nodes. If you are looking for stake your funds, you can take a closer look at my options.

Nymesis Explorer: https://nymesis.vercel.app/?q=3826
