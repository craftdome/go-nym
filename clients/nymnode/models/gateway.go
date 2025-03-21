package models

type MixnetWebsockets struct {
	WsPort  uint16 `json:"ws_port"`
	WssPort uint16 `json:"wss_port,omitempty"`
}

type Wireguard struct {
	Port      uint16 `json:"port"`
	PublicKey string `json:"public_key"`
}

type ClientInterfaces struct {
	MixnetWebsockets `json:"mixnet_websockets,omitempty"`
	Wireguard        `json:"wireguard,omitempty"`
}

type Gateway struct {
	ClientInterfaces `json:"client_interfaces,omitempty"`

	EnforcesZKNyms bool `json:"enforces_zk_nyms"`
}
