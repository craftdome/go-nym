package nym

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"os"
	"time"
)

type Client struct {
	server string
	conn   *websocket.Conn

	messages chan Response
}

func NewClient(server string) *Client {
	return &Client{
		server:   server,
		messages: make(chan Response, 1),
	}
}

func (c *Client) Dial() (err error) {
	c.conn, _, err = websocket.DefaultDialer.Dial(c.server, nil)
	if err != nil {
		return errors.Wrapf(err, "websocket.DefaultDialer.Dial %s", c.server)
	}

	return nil
}

func (c *Client) ListenAndServe() error {
	defer close(c.messages)
	for {
		messageType, r, err := c.conn.NextReader()
		if err != nil {
			return err
		}

		switch messageType {
		case websocket.TextMessage:
			packet, err := FromJSON(r)
			if err != nil {
				fmt.Fprintln(os.Stderr, errors.Wrap(err, ""))
				continue
			}

			c.messages <- packet
		case websocket.BinaryMessage:
			fmt.Fprintln(os.Stderr, ErrBinaryMessageNotSupported)
		}
	}
}

func (c *Client) Close() error {
	if c.conn != nil {
		message := websocket.FormatCloseMessage(websocket.CloseMessage, "")
		if err := c.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(10*time.Second)); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Messages() <-chan Response {
	return c.messages
}

func (c *Client) SendRequestAsText(r Request) error {
	if c.conn != nil {
		data, err := r.ToJSON()
		if err != nil {
			return errors.Wrap(err, "ToJSON")
		}
		if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return errors.Wrapf(err, "conn.WriteMessage(websocket.TextMessage, %s)", string(data))
		}

		return nil
	}

	return ErrConnectionNotEstablished
}

func (c *Client) SendRequestAsBinary(r Request) error {
	panic("implement!")
}
