package intervals

import (
	"context"
	"github.com/craftdome/go-nym/pkg/mixnet"
)

func (c *Client) GetPendingEpochEvents(ctx context.Context, params GetPendingEpochEventsParams) (mixnet.PendingEpochEvents, error) {
	type req struct {
		MethodParams GetPendingEpochEventsParams `json:"get_pending_epoch_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.EpochEventsMaxRetrievalLimit)

	return Query[mixnet.PendingEpochEvents](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingIntervalEvents(ctx context.Context, params GetPendingIntervalEventsParams) (mixnet.PendingIntervalEvents, error) {
	type req struct {
		MethodParams GetPendingIntervalEventsParams `json:"get_pending_interval_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.IntervalEventsMaxRetrievalLimit)

	return Query[mixnet.PendingIntervalEvents](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingEpochEvent(ctx context.Context, params GetPendingEpochEventParams) (mixnet.PendingEpochEvent, error) {
	type req struct {
		MethodParams GetPendingEpochEventParams `json:"get_pending_epoch_event"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.PendingEpochEvent](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingIntervalEvent(ctx context.Context, params GetPendingIntervalEventParams) (mixnet.PendingIntervalEvent, error) {
	type req struct {
		MethodParams GetPendingIntervalEventParams `json:"get_pending_interval_event"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.PendingIntervalEvent](ctx, c.client, c.contract, r)
}

func (c *Client) GetNumberOfPendingEvents(ctx context.Context) (mixnet.NumberOfPendingEvents, error) {
	type req struct {
		MethodParams struct{} `json:"get_number_of_pending_events"`
	}

	r := req{}

	return Query[mixnet.NumberOfPendingEvents](ctx, c.client, c.contract, r)
}
