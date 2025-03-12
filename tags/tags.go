package tags

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var ErrUnknownResponseTag = errors.New("unknown response tag")

// Requests
var (
	Send               RequestTag = &sendType{requestCmd{"send", 0x0}}
	SendAnonymous      RequestTag = &sendAnonymousType{requestCmd{"sendAnonymous", 0x1}}
	Reply              RequestTag = &replyType{requestCmd{"reply", 0x2}}
	GetSelfAddress     RequestTag = &getSelfAddressType{requestCmd{"selfAddress", 0x3}}
	ClosedConnection   RequestTag = &closedConnectionType{requestCmd{"closedConnection", 0x4}}
	GetLaneQueueLength RequestTag = &getLaneQueueLengthType{requestCmd{"getLaneQueueLength", 0x5}}
)

// Responses
var (
	Error           ResponseTag = &errorType{responseCmd{"error", 0x0}}
	Received        ResponseTag = &receivedType{responseCmd{"received", 0x1}}
	SelfAddress     ResponseTag = &selfAddressType{responseCmd{"selfAddress", 0x2}}
	LaneQueueLength ResponseTag = &laneQueueLengthType{responseCmd{"laneQueueLength", 0x3}}
)

type RequestTag interface {
	Text() string
	Binary() byte
}

type requestCmd struct {
	text   string
	binary byte
}

func (c *requestCmd) Text() string {
	return c.text
}

func (c *requestCmd) Binary() byte {
	return c.binary
}

type sendType struct {
	requestCmd
}

type sendAnonymousType struct {
	requestCmd
}

type replyType struct {
	requestCmd
}

type getSelfAddressType struct {
	requestCmd
}

type closedConnectionType struct {
	requestCmd
}

type getLaneQueueLengthType struct {
	requestCmd
}

type ResponseTag interface {
	Text() string
	Binary() byte
}

type responseCmd struct {
	text   string
	binary byte
}

func (c *responseCmd) Text() string {
	return c.text
}

func (c *responseCmd) Binary() byte {
	return c.binary
}

type errorType struct {
	responseCmd
}

type receivedType struct {
	responseCmd
}

type selfAddressType struct {
	responseCmd
}

type laneQueueLengthType struct {
	responseCmd
}

type responseType struct {
	Type string `json:"type"`
}

func GetResponseTagFromJSON(data []byte) (ResponseTag, error) {
	var t responseType
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	switch t.Type {
	case Error.Text():
		return Error, nil
	case Received.Text():
		return Received, nil
	case SelfAddress.Text():
		return SelfAddress, nil
	case LaneQueueLength.Text():
		return LaneQueueLength, nil
	default:
		return nil, errors.Wrap(ErrUnknownResponseTag, t.Type)
	}
}

func GetResponseTagFromBinary(firstByte byte) (ResponseTag, error) {
	switch firstByte {
	case Error.Binary():
		return Error, nil
	case Received.Binary():
		return Received, nil
	case SelfAddress.Binary():
		return SelfAddress, nil
	case LaneQueueLength.Binary():
		return LaneQueueLength, nil
	default:
		return nil, errors.Wrapf(ErrUnknownResponseTag, "0x%X", firstByte)
	}
}
