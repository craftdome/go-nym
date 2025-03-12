package wsc

import (
	request2 "github.com/craftdome/go-nym/wsc/request"
	response2 "github.com/craftdome/go-nym/wsc/response"
	"github.com/craftdome/go-nym/wsc/tags"
	"github.com/pkg/errors"
	"io"
)

type Request interface {
	ToJSON() ([]byte, error)
	//Write(writer io.WriteCloser) error
}

func NewSend(message, recipient string) Request {
	return &request2.Send{
		Message:   message,
		Recipient: recipient,
	}
}

func NewSendAnonymous(message, recipient string, replySurbs int) Request {
	return &request2.SendAnonymous{
		Message:    message,
		Recipient:  recipient,
		ReplySurbs: replySurbs,
	}
}

func NewReply(message, senderTag string) Request {
	return &request2.Reply{
		Message:   message,
		SenderTag: senderTag,
	}
}

func NewGetSelfAddress() Request {
	return &request2.GetSelfAddress{}
}

func NewClosedConnection() Request {
	return &request2.ClosedConnection{}
}

func NewGetLaneQueueLength() Request {
	return &request2.GetLaneQueueLength{}
}

type Response interface {
	Type() tags.ResponseTag
}

func FromJSON(reader io.Reader) (Response, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "io.ReadAll")
	}

	tag, err := tags.GetResponseTagFromJSON(data)
	if err != nil {
		return nil, errors.Wrap(err, "tags.GetResponseTagFromJSON")
	}

	switch tag {
	case tags.Error:
		r := &response2.Error{}
		if err := r.InitFromJSON(data); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromJSON data to %s", tag.Text())
		}
		return r, nil
	case tags.Received:
		r := &response2.Received{}
		if err := r.InitFromJSON(data); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromJSON data to %s", tag.Text())
		}
		return r, nil
	case tags.SelfAddress:
		r := &response2.SelfAddress{}
		if err := r.InitFromJSON(data); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromJSON data to %s", tag.Text())
		}
		return r, nil
	case tags.LaneQueueLength:
		r := &response2.LaneQueueLength{}
		if err := r.InitFromJSON(data); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromJSON data to %s", tag.Text())
		}
		return r, nil
	default:
		return nil, errors.New("does not correspond to any valid request")
	}
}

func FromBinary(reader io.Reader) (Response, error) {
	var firstByte [1]byte
	if _, err := reader.Read(firstByte[:]); err != nil {
		return nil, errors.Wrap(err, "read first byte")
	}

	tag, err := tags.GetResponseTagFromBinary(firstByte[0])
	if err != nil {
		return nil, errors.Wrapf(err, "tags.GetResponseTagFromBinary(0x%X)", firstByte[0])
	}

	switch tag {
	case tags.Error:
		r := &response2.Error{}
		if err := r.InitFromBinary(reader); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromBinary data to %s", tag.Text())
		}
		return r, nil
	case tags.Received:
		r := &response2.Received{}
		if err := r.InitFromBinary(reader); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromBinary data to %s", tag.Text())
		}
		return r, nil
	case tags.SelfAddress:
		r := &response2.SelfAddress{}
		if err := r.InitFromBinary(reader); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromBinary data to %s", tag.Text())
		}
		return r, nil
	case tags.LaneQueueLength:
		r := &response2.LaneQueueLength{}
		if err := r.InitFromBinary(reader); err != nil {
			return nil, errors.Wrapf(err, "cannot InitFromBinary data to %s", tag.Text())
		}
		return r, nil
	default:
		return nil, errors.New("does not correspond to any valid request")
	}
}
