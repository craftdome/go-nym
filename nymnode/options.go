package nymnode

import (
	"golang.org/x/time/rate"
	"net/http"
)

type Option func(*Client)

func WithCustomClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

func WithRateLimiter(limiter *rate.Limiter) Option {
	return func(c *Client) {
		c.limiter = limiter
	}
}
