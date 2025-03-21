package nymapi_test

import (
	"github.com/craftdome/go-nym/clients/nymapi"
	"testing"
)

const (
	host = "validator.nymtech.net"
)

func setup(t *testing.T) *nymapi.Client {
	c, err := nymapi.New(t.Context(), host)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

func teardown(t *testing.T, c *nymapi.Client) {}

func TestClient(t *testing.T) {
	c := setup(t)
	t.Cleanup(func() { teardown(t, c) })

	t.Run("GetBuildInformation", func(t *testing.T) {
		t.Parallel()
		res, err := c.GetBuildInformation(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	})

	t.Run("Health", func(t *testing.T) {
		t.Parallel()
		res, err := c.Health(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	})

	t.Run("GetNodeAnnotation", func(t *testing.T) {
		t.Parallel()
		res, err := c.GetNodeAnnotation(t.Context(), 895)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	})

	t.Run("GetNodePerformance", func(t *testing.T) {
		t.Parallel()
		res, err := c.GetNodePerformance(t.Context(), 895)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	})
}
