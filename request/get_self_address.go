package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/tags"
)

type GetSelfAddress struct {
}

func (r *GetSelfAddress) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": tags.GetSelfAddress.Text(),
	})
}
