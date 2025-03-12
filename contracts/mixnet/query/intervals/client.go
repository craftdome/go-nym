package intervals

import (
	"contracts/mixnet/shared"
)

type Client struct {
	reader shared.ContractReader
}

func New(reader shared.ContractReader) *Client {
	return &Client{
		reader: reader,
	}
}
