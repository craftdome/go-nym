package api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
)

func get[T comparable](ctx context.Context, client *http.Client, limiter *rate.Limiter, url string, result T) error {
	if err := limiter.Wait(ctx); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
