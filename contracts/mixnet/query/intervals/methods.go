package intervals

import (
	"context"
	"github.com/craftdome/go-nym/contracts/mixnet/models"
)

func (c *Client) GetPendingEpochEvents(ctx context.Context, params GetPendingEpochEventsParams) (*models.PendingEpochEvents, error) {
	type req struct {
		MethodParams GetPendingEpochEventsParams `json:"get_pending_epoch_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.EpochEventsMaxRetrievalLimit)

	var result models.PendingEpochEvents
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetPendingIntervalEvents(ctx context.Context, params GetPendingIntervalEventsParams) (*models.PendingIntervalEvents, error) {
	type req struct {
		MethodParams GetPendingIntervalEventsParams `json:"get_pending_interval_events"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.IntervalEventsMaxRetrievalLimit)

	var result models.PendingIntervalEvents
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetPendingEpochEvent(ctx context.Context, params GetPendingEpochEventParams) (*models.PendingEpochEvent, error) {
	type req struct {
		MethodParams GetPendingEpochEventParams `json:"get_pending_epoch_event"`
	}

	r := req{MethodParams: params}

	var result models.PendingEpochEvent
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetPendingIntervalEvent(ctx context.Context, params GetPendingIntervalEventParams) (*models.PendingIntervalEvent, error) {
	type req struct {
		MethodParams GetPendingIntervalEventParams `json:"get_pending_interval_event"`
	}

	r := req{MethodParams: params}

	var result models.PendingIntervalEvent
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetNumberOfPendingEvents(ctx context.Context) (*models.NumberOfPendingEvents, error) {
	type req struct {
		MethodParams struct{} `json:"get_number_of_pending_events"`
	}

	r := req{}

	var result models.NumberOfPendingEvents
	return &result, c.reader.Read(ctx, r, &result)
}
