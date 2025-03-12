package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/wsc/tags"
)

type ClosedConnection struct {
}

func (r *ClosedConnection) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": tags.ClosedConnection.Text(),
	})
}
