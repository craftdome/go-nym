package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/tags"
)

type Send struct {
	Message   string
	Recipient string
}

func (r *Send) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":      tags.Send.Text(),
		"message":   r.Message,
		"recipient": r.Recipient,
	})
}
