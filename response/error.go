package response

import (
	"encoding/json"
	"github.com/Tyz3/go-nym/tags"
	"io"
)

type Error struct {
	Message string `json:"message"`
}

func (r *Error) Type() tags.ResponseTag {
	return tags.Error
}

func (r *Error) InitFromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Error) InitFromBinary(reader io.Reader) error {
	// TODO
	panic("implement!")
}
