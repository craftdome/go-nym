package nymnode

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"io"
	"net/http"
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

func (c *Client) GetAuxiliaryDetails(ctx context.Context) (AuxiliaryDetails, error) {
	return get[AuxiliaryDetails](ctx, c.client, c.limiter, c.endpoints[EndpointAuxiliaryDetails])
}

func (c *Client) GetBuildInformation(ctx context.Context) (BuildInformation, error) {
	return get[BuildInformation](ctx, c.client, c.limiter, c.endpoints[EndpointBuildInformation])
}

func (c *Client) GetDescription(ctx context.Context) (Description, error) {
	return get[Description](ctx, c.client, c.limiter, c.endpoints[EndpointDescription])
}

func (c *Client) GetHostInformation(ctx context.Context) (HostInformation, error) {
	return get[HostInformation](ctx, c.client, c.limiter, c.endpoints[EndpointHostInformation])
}

func (c *Client) GetRoles(ctx context.Context) (Roles, error) {
	return get[Roles](ctx, c.client, c.limiter, c.endpoints[EndpointRoles])
}

func (c *Client) GetSystemInformation(ctx context.Context) (SystemInformation, error) {
	return get[SystemInformation](ctx, c.client, c.limiter, c.endpoints[EndpointSystemInformation])
}

func (c *Client) GetGateway(ctx context.Context) (Gateway, error) {
	return get[Gateway](ctx, c.client, c.limiter, c.endpoints[EndpointGateway])
}

func (c *Client) GetGatewayClientInterfaces(ctx context.Context) (ClientInterfaces, error) {
	return get[ClientInterfaces](ctx, c.client, c.limiter, c.endpoints[EndpointGatewayClientInterfaces])
}

func (c *Client) GetGatewayClientInterfacesMixnetWebsockets(ctx context.Context) (MixnetWebsockets, error) {
	return get[MixnetWebsockets](ctx, c.client, c.limiter, c.endpoints[EndpointGatewayClientInterfacesMixnetWebsockets])
}

func (c *Client) Health(ctx context.Context) (Health, error) {
	return get[Health](ctx, c.client, c.limiter, c.endpoints[EndpointHealth])
}

func (c *Client) GetIPR(ctx context.Context) (IPPacketRouter, error) {
	return get[IPPacketRouter](ctx, c.client, c.limiter, c.endpoints[EndpointIPR])
}

func (c *Client) GetPacketStatsMetrics(ctx context.Context) (PacketsStatsMetrics, error) {
	return get[PacketsStatsMetrics](ctx, c.client, c.limiter, c.endpoints[EndpointMetricsPacketsStats])
}

func (c *Client) GetPrometheusMetrics(ctx context.Context, token string) (PrometheusMetrics, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoints[EndpointMetricsPrometheus], nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return PrometheusMetrics(data), nil
}

func (c *Client) GetNR(ctx context.Context) (NetworkRequester, error) {
	return get[NetworkRequester](ctx, c.client, c.limiter, c.endpoints[EndpointNR])
}

func (c *Client) GetNRExitPolicy(ctx context.Context) (NetworkRequesterExitPolicy, error) {
	return get[NetworkRequesterExitPolicy](ctx, c.client, c.limiter, c.endpoints[EndpointNRExitPolicy])
}
