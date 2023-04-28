package request

import (
	"encoding/json"
	"github.com/Tyz3/go-nym/tags"
)

type ClosedConnection struct {
}

func (r *ClosedConnection) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": tags.ClosedConnection.Text(),
	})
}
