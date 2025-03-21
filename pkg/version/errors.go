package version

import "errors"

var (
	ErrInsufficientParts = errors.New("insufficient parts, must be x.y.z")
)
