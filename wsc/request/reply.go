package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/wsc/tags"
)

type Reply struct {
	Message   string
	SenderTag string
}

func (r *Reply) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":      tags.Reply.Text(),
		"message":   r.Message,
		"senderTag": r.SenderTag,
	})
}
