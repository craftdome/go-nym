![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/craftdome/go-nym?style=flat-square)
![GitHub release (with filter)](https://img.shields.io/github/v/release/craftdome/go-nym?style=flat-square)

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

Use the standard Go tools to install dependencies:
```
go get github.com/craftdome/go-nym
```

Importing the basic package:
```go
import "github.com/craftdome/go-nym"
```

# Using

### Tips

- Remember, if you plan to give access to `nym-client` connect from the outside (for this you should specify the ip of the external network interface of your machine), the `nym-client` doesn't have a connection auth function. 
- Only 1 connection is allowed at a time.
- If you need an external connection, use a local networks ([10.0.0.0/8, 192.168.0.0/16, 172.16.0.0/12](https://en.wikipedia.org/wiki/Private_network)) instead of a global unicast one to increase security.

### Initialization

1. First, you need to `init` the connection client with the nym-client address. We copy the connection `address:port` from the console after running the nym-client. By default it's `localhost:1977`, or in my case `192.168.88.4:1977`.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L21-L22

2. Connection establishing with the `nym-client`.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L24-L27

### Reading messages

3. We turn on listening to incoming messages, which we then extract through the `Messages()` chan.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L30-L54

### Sending messages

4. Getting your nym-client address (SelfAddress).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L57-L59

5. Sending a message (Send).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L61-L66

6. Sending a SURB message to receive an anonymous response (SendAnonymous).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L68-L74

7. Sending a reply to the SendAnonymous message (Reply).

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L76-L81

### Closing the connection correctly

8. We close the connection with the `nym-client` after the interrupt (sigint/sigkill) and wait for the reading gorutine sent `done` signal.

https://github.com/craftdome/go-nym/blob/17a1c03713c834ca65295c0c5818b0d1dba364e2/example/main.go#L83-L91

# Support the developer (Mixnodes)

Below is a list of the developer's mixnodes. If you are looking for a node for delegating tokens, you can take a closer look at my options.

[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.identity_key&style=flat-square&logo=numpy&logoColor=white&label=Advanced%20Engineering%201&color=%23136401&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=stake_saturation&style=flat-square&logo=myspace&logoColor=white&label=Stake&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=profit_margin_percent&style=flat-square&logo=buymeacoffee&logoColor=white&label=Owner%20Profit)](https://explorer.nymtech.net/network-components/mixnode/895)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F895&query=mix_node.version&style=flat-square&logo=git&logoColor=white&label=Version&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/895)

[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=mix_node.identity_key&style=flat-square&logo=numpy&logoColor=white&label=Advanced%20Engineering%202&color=%23136401&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=stake_saturation&style=flat-square&logo=myspace&logoColor=white&label=Stake&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=profit_margin_percent&style=flat-square&logo=buymeacoffee&logoColor=white&label=Owner%20Profit)](https://explorer.nymtech.net/network-components/mixnode/1227)
[![Dynamic JSON Badge](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fexplorer.nymtech.net%2Fapi%2Fv1%2Fmix-node%2F1227&query=mix_node.version&style=flat-square&logo=git&logoColor=white&label=Version&cacheSeconds=60)](https://explorer.nymtech.net/network-components/mixnode/1227)
