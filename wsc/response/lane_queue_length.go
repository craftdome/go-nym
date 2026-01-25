package response

import (
	"encoding/json"
	"io"

	"github.com/craftdome/go-nym/wsc/tags"
)

type LaneQueueLength struct {
}

func (r *LaneQueueLength) Type() tags.ResponseTag {
	return tags.LaneQueueLength
}

func (r *LaneQueueLength) InitFromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *LaneQueueLength) InitFromBinary(reader io.Reader) error {
	// TODO
	panic("implement!")
}
