package request

import (
	"encoding/json"
	"github.com/craftdome/go-nym/tags"
)

type GetLaneQueueLength struct {
}

func (r *GetLaneQueueLength) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type": tags.GetLaneQueueLength.Text(),
	})
}
