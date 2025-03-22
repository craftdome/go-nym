package nymapi

import (
	"context"
	"encoding/json"
	"github.com/craftdome/go-nym/pkg/mixnet"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"net/http"
	"strings"
)

func get[T any](ctx context.Context, client *http.Client, limiter *rate.Limiter, url string) (result T, err error) {
	if err = limiter.Wait(ctx); err != nil {
		return result, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, errors.Errorf("%s -> %s", url, resp.Status)
	}

	return result, json.NewDecoder(resp.Body).Decode(&result)
}

func (c *Client) GetBuildInformation(ctx context.Context) (BuildInformation, error) {
	return get[BuildInformation](ctx, c.client, c.limiter, c.endpoints[EndpointBuildInformation])
}

func (c *Client) Health(ctx context.Context) (Health, error) {
	return get[Health](ctx, c.client, c.limiter, c.endpoints[EndpointHealth])
}

func (c *Client) GetNodeAnnotation(ctx context.Context, nodeID mixnet.NodeID) (NodeAnnotation, error) {
	url := strings.Replace(c.endpoints[EndpointBuildInformation], "{node_id}", nodeID.String(), 1)
	return get[NodeAnnotation](ctx, c.client, c.limiter, url)
}

func (c *Client) GetNodePerformance(ctx context.Context, nodeID mixnet.NodeID) (NodePerformance, error) {
	url := strings.Replace(c.endpoints[EndpointNodePerformance], "{node_id}", nodeID.String(), 1)
	return get[NodePerformance](ctx, c.client, c.limiter, url)
}
