package nymapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"

	"github.com/craftdome/go-nym/pkg/version"
)

// Minimum supported API Versions
var (
	v1154 = version.MustParse("1.1.54")
)

var ErrUnsupportedNodeVersion = errors.New("unsupported node version")

type Client struct {
	client  *http.Client
	limiter *rate.Limiter

	endpoints map[string]string

	BuildInformation BuildInformation
}

func New(ctx context.Context, host string, opts ...Option) (*Client, error) {
	c := &Client{
		client:    &http.Client{},
		limiter:   rate.NewLimiter(rate.Inf, 0),
		endpoints: make(map[string]string, len(Endpoints)),
	}

	for _, endpoint := range Endpoints {
		c.endpoints[endpoint] = fmt.Sprintf(EndpointTemplate, host, endpoint)
	}

	for _, opt := range opts {
		opt(c)
	}

	info, err := c.getBuildInformation(ctx)
	if err != nil {
		return nil, err
	}

	c.BuildInformation = info

	v, err := version.Parse(info.BuildVersion)
	if err != nil {
		return nil, err
	}

	if v < v1154 {
		return nil, errors.Wrapf(ErrUnsupportedNodeVersion, "%s < %s", v, v1154)
	}

	return c, nil
}

func (c *Client) getBuildInformation(ctx context.Context) (BuildInformation, error) {
	return get[BuildInformation](ctx, c.client, c.limiter, c.endpoints[EndpointBuildInformation])
}
