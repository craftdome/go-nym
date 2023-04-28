package nym

import "github.com/pkg/errors"

var (
	ErrBinaryMessageNotSupported = errors.New("websocket frame BinaryMessage not supported yet")
	ErrConnectionNotEstablished  = errors.New("connection is not established")
)
