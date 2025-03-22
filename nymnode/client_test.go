package nymnode_test

import (
	"github.com/craftdome/go-nym/nymnode"
	"testing"
)

const (
	host = "38.180.189.128:8080" // "38.180.189.128", "2001:ac8:a:27:0:2:0:421"
)

func setup(t *testing.T) *nymnode.Client {
	c, err := nymnode.New(t.Context(), host)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

func teardown(t *testing.T, c *nymnode.Client) {}

func TestClient(t *testing.T) {
	c := setup(t)
	t.Cleanup(func() { teardown(t, c) })

	t.Run("Node GetAuxiliaryDetails", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetAuxiliaryDetails(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Node GetBuildInformation", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetBuildInformation(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Node GetDescription", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetDescription(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Node GetHostInformation", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetHostInformation(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Node GetRoles", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetRoles(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Node GetSystemInformation", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetSystemInformation(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Gateway GetGateway", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetGateway(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Gateway GetGatewayClientInterfaces", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetGatewayClientInterfaces(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Gateway GetGatewayClientInterfacesMixnetWebsockets", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetGatewayClientInterfacesMixnetWebsockets(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Health Health", func(t *testing.T) {
		t.Parallel()
		info, err := c.Health(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("IPR GetIPR", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetIPR(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("Metrics GetPacketStatsMetrics", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetPacketStatsMetrics(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("NR GetNR", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetNR(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})

	t.Run("NR GetNRExitPolicy", func(t *testing.T) {
		t.Parallel()
		info, err := c.GetNRExitPolicy(t.Context())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v", info)
	})
}
