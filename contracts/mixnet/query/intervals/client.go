package intervals

import (
	"github.com/craftdome/go-nym/contracts/v1/mixnet/shared"
)

type Client struct {
	reader shared.ContractReader
}

func New(reader shared.ContractReader) *Client {
	return &Client{
		reader: reader,
	}
}
