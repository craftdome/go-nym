package api

import (
	"context"
	"fmt"
	"github.com/craftdome/go-nym/clients/nymnode/models"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"io"
	"net/http"
)

type Client131 struct {
	client  *http.Client
	limiter *rate.Limiter

	endpoints map[string]string
}

func New131(client *http.Client, limiter *rate.Limiter, host string, port uint16) *Client131 {
	c := &Client131{
		client:    client,
		limiter:   limiter,
		endpoints: make(map[string]string, len(Endpoints)),
	}

	for _, endpoint := range Endpoints {
		c.endpoints[endpoint] = fmt.Sprintf(EndpointTemplate, host, port, endpoint)
	}

	return c
}

func (c *Client131) GetAuxiliaryDetails(ctx context.Context) (models.AuxiliaryDetails, error) {
	var result models.AuxiliaryDetails

	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointAuxiliaryDetails], &result)
}

func (c *Client131) GetBuildInformation(ctx context.Context) (models.BuildInformation, error) {
	var result models.BuildInformation
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointBuildInformation], &result)
}

func (c *Client131) GetDescription(ctx context.Context) (models.Description, error) {
	var result models.Description
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointDescription], &result)
}

func (c *Client131) GetHostInformation(ctx context.Context) (models.HostInformation, error) {
	var result models.HostInformation
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointHostInformation], &result)
}

func (c *Client131) GetRoles(ctx context.Context) (models.Roles, error) {
	var result models.Roles
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointRoles], &result)
}

func (c *Client131) GetSystemInformation(ctx context.Context) (models.SystemInformation, error) {
	var result models.SystemInformation
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointSystemInformation], &result)
}

func (c *Client131) GetGateway(ctx context.Context) (models.Gateway, error) {
	var result models.Gateway
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointGateway], &result)
}

func (c *Client131) GetGatewayClientInterfaces(ctx context.Context) (models.ClientInterfaces, error) {
	var result models.ClientInterfaces
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointGatewayClientInterfaces], &result)
}

func (c *Client131) GetGatewayClientInterfacesMixnetWebsockets(ctx context.Context) (models.MixnetWebsockets, error) {
	var result models.MixnetWebsockets
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointGatewayClientInterfacesMixnetWebsockets], &result)
}

func (c *Client131) Health(ctx context.Context) (models.Health, error) {
	var result models.Health
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointHealth], &result)
}

func (c *Client131) GetIPR(ctx context.Context) (models.IPPacketRouter, error) {
	var result models.IPPacketRouter
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointIPR], &result)
}

func (c *Client131) GetPacketStatsMetrics(ctx context.Context) (models.PacketsStatsMetrics, error) {
	var result models.PacketsStatsMetrics
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointMetricsPacketsStats], &result)
}

func (c *Client131) GetPrometheusMetrics(ctx context.Context, token string) (models.PrometheusMetrics, error) {
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

	return models.PrometheusMetrics(data), nil
}

func (c *Client131) GetNR(ctx context.Context) (models.NetworkRequester, error) {
	var result models.NetworkRequester
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointNR], &result)
}

func (c *Client131) GetNRExitPolicy(ctx context.Context) (models.NetworkRequesterExitPolicy, error) {
	var result models.NetworkRequesterExitPolicy
	return result, get(ctx, c.client, c.limiter, c.endpoints[EndpointNRExitPolicy], &result)
}
