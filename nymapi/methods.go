package nymapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"

	"github.com/craftdome/go-nym/pkg/types"
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

func (c *Client) Health(ctx context.Context) (Health, error) {
	return get[Health](ctx, c.client, c.limiter, c.endpoints[EndpointHealth])
}

func (c *Client) GetNodeAnnotation(ctx context.Context, nodeID types.NodeID) (NodeAnnotation, error) {
	url := strings.Replace(c.endpoints[EndpointNodeAnnotation], "{node_id}",
		strconv.FormatUint(uint64(nodeID), 10),
		1,
	)
	return get[NodeAnnotation](ctx, c.client, c.limiter, url)
}

func (c *Client) GetNodePerformance(ctx context.Context, nodeID types.NodeID) (NodePerformance, error) {
	url := strings.Replace(c.endpoints[EndpointNodePerformance], "{node_id}",
		strconv.FormatUint(uint64(nodeID), 10),
		1,
	)
	return get[NodePerformance](ctx, c.client, c.limiter, url)
}
