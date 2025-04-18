package nymnode

import (
	"context"
	"fmt"
	"github.com/craftdome/go-nym/pkg/version"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"net/http"
)

// Minimum supported API Versions
var (
	v131 = version.MustParse("1.3.1")
	v150 = version.MustParse("1.5.0")
)

var ErrUnsupportedNodeVersion = errors.New("unsupported node version")

type Client struct {
	client  *http.Client
	limiter *rate.Limiter

	endpoints map[string]string

	BuildInformation BuildInformation
	Version          version.Version
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

	if v < v131 {
		return nil, errors.Wrapf(ErrUnsupportedNodeVersion, "current=%s < %s", v, v131)
	}

	c.Version = v

	return c, nil
}

func (c *Client) getBuildInformation(ctx context.Context) (BuildInformation, error) {
	return get[BuildInformation](ctx, c.client, c.limiter, c.endpoints[EndpointBuildInformation])
}
