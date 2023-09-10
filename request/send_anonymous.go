package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/tags"
)

type SendAnonymous struct {
	Message    string
	Recipient  string
	ReplySurbs int
}

func (r *SendAnonymous) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":       tags.SendAnonymous.Text(),
		"message":    r.Message,
		"recipient":  r.Recipient,
		"replySurbs": r.ReplySurbs,
	})
}
