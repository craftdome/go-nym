package intervals

import (
	"github.com/craftdome/go-nym/pkg/mixnet"
)

type GetPendingEpochEventsParams struct {
	StartAfter mixnet.EpochEventID `json:"start_after,omitzero"`
	Limit      uint32              `json:"limit,omitempty"`
}

type GetPendingIntervalEventsParams struct {
	StartAfter mixnet.IntervalEventID `json:"start_after,omitzero"`
	Limit      uint32                 `json:"limit,omitempty"`
}

type GetPendingEpochEventParams struct {
	EventID mixnet.EpochEventID `json:"event_id"`
}

type GetPendingIntervalEventParams struct {
	EventID mixnet.IntervalEventID `json:"event_id"`
}
