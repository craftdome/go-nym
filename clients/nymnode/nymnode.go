package nymnode

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/craftdome/go-nym/clients/nymnode/api"
	"github.com/craftdome/go-nym/clients/nymnode/models"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"net/http"
	"version"
)

// Supported API Versions
var (
	v131 = version.MustParse("1.3.1")
)

type Node interface {
	GetAuxiliaryDetails(ctx context.Context) (models.AuxiliaryDetails, error)
	GetBuildInformation(ctx context.Context) (models.BuildInformation, error)
	GetDescription(ctx context.Context) (models.Description, error)
	GetHostInformation(ctx context.Context) (models.HostInformation, error)
	GetRoles(ctx context.Context) (models.Roles, error)
	GetSystemInformation(ctx context.Context) (models.SystemInformation, error)
}

type Gateway interface {
	GetGateway(ctx context.Context) (models.Gateway, error)
	GetGatewayClientInterfaces(ctx context.Context) (models.ClientInterfaces, error)
	GetGatewayClientInterfacesMixnetWebsockets(ctx context.Context) (models.MixnetWebsockets, error)
}

type Health interface {
	Health(ctx context.Context) (models.Health, error)
}

type IPR interface {
	GetIPR(ctx context.Context) (models.IPPacketRouter, error)
}

type Metrics interface {
	GetPacketStatsMetrics(ctx context.Context) (models.PacketsStatsMetrics, error)
	GetPrometheusMetrics(ctx context.Context, token string) (models.PrometheusMetrics, error)
}

type NR interface {
	GetNR(ctx context.Context) (models.NetworkRequester, error)
	GetNRExitPolicy(ctx context.Context) (models.NetworkRequesterExitPolicy, error)
}

type Client interface {
	Node
	Gateway
	Health
	IPR
	Metrics
	NR
}

func New(ctx context.Context, client *http.Client, limiter *rate.Limiter, host string, port uint16) (Client, error) {
	info, err := getBuildInformation(ctx, client, host, port)
	if err != nil {
		return nil, err
	}

	v, err := version.Parse(info.BuildVersion)
	if err != nil {
		return nil, err
	}

	switch {
	case v >= v131:
		return api.New131(client, limiter, host, port), nil
	default:
		return nil, fmt.Errorf("unsupported version")
	}
}

func getBuildInformation(ctx context.Context, client *http.Client, host string, port uint16) (*models.BuildInformation, error) {
	url := fmt.Sprintf("http://%s:%d/api/v1/build-information", host, port)

	var result models.BuildInformation
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return &result, json.NewDecoder(resp.Body).Decode(&result)
}
