package nymnode

import "github.com/pkg/errors"

var (
	ErrUnsupportedNodeVersion = errors.New("unsupported node version")
	ErrUnsupportedNodeMethod  = errors.New("unsupported node method")

	ErrForbiddenNodeMethod = errors.New("forbidden to access node method")
)
