package models

type IPPacketRouter struct {
	EncodedIdentityKey string `json:"encoded_identity_key"`
	EncodedX25519Key   string `json:"encoded_x25519_key"`
	Address            string `json:"address"`
}
