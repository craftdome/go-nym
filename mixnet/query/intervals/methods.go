package intervals

import (
	"context"

	"github.com/craftdome/go-nym/pkg/types"
)

func (c *Client) GetPendingEpochEvents(ctx context.Context, params GetPendingEpochEventsParams) (types.PendingEpochEvents, error) {
	type req struct {
		MethodParams GetPendingEpochEventsParams `json:"get_pending_epoch_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.EpochEventsMaxRetrievalLimit)

	return Query[types.PendingEpochEvents](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingIntervalEvents(ctx context.Context, params GetPendingIntervalEventsParams) (types.PendingIntervalEvents, error) {
	type req struct {
		MethodParams GetPendingIntervalEventsParams `json:"get_pending_interval_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.IntervalEventsMaxRetrievalLimit)

	return Query[types.PendingIntervalEvents](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingEpochEvent(ctx context.Context, params GetPendingEpochEventParams) (types.PendingEpochEvent, error) {
	type req struct {
		MethodParams GetPendingEpochEventParams `json:"get_pending_epoch_event"`
	}

	r := req{MethodParams: params}

	return Query[types.PendingEpochEvent](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingIntervalEvent(ctx context.Context, params GetPendingIntervalEventParams) (types.PendingIntervalEvent, error) {
	type req struct {
		MethodParams GetPendingIntervalEventParams `json:"get_pending_interval_event"`
	}

	r := req{MethodParams: params}

	return Query[types.PendingIntervalEvent](ctx, c.client, c.contract, r)
}

func (c *Client) GetNumberOfPendingEvents(ctx context.Context) (types.NumberOfPendingEvents, error) {
	type req struct {
		MethodParams struct{} `json:"get_number_of_pending_events"`
	}

	r := req{}

	return Query[types.NumberOfPendingEvents](ctx, c.client, c.contract, r)
}
