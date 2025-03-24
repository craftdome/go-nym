package intervals

import (
	"github.com/craftdome/go-nym/pkg/types"
)

type GetPendingEpochEventsParams struct {
	StartAfter types.EpochEventID `json:"start_after,omitempty"`
	Limit      uint32             `json:"limit,omitempty"`
}

type GetPendingIntervalEventsParams struct {
	StartAfter types.IntervalEventID `json:"start_after,omitempty"`
	Limit      uint32                `json:"limit,omitempty"`
}

type GetPendingEpochEventParams struct {
	EventID types.EpochEventID `json:"event_id"`
}

type GetPendingIntervalEventParams struct {
	EventID types.IntervalEventID `json:"event_id"`
}
