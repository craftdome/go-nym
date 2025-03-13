package response

import (
	"encoding/json"
	"github.com/craftdome/go-nym/clients/wsc/tags"
	"io"
)

type SelfAddress struct {
	Address string `json:"address"`
}

func (r *SelfAddress) Type() tags.ResponseTag {
	return tags.SelfAddress
}

func (r *SelfAddress) InitFromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *SelfAddress) InitFromBinary(reader io.Reader) error {
	// TODO
	panic("implement!")
}
