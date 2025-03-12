package response

import (
	"encoding/json"
	"github.com/craftdome/go-nym/wsc/tags"
	"io"
)

type Received struct {
	Message   string `json:"message"`
	SenderTag string `json:"senderTag,omitempty"`
}

func (r *Received) Type() tags.ResponseTag {
	return tags.Received
}

func (r *Received) InitFromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Received) InitFromBinary(reader io.Reader) error {
	// TODO
	panic("implement!")
}
