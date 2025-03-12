package intervals

import (
	"github.com/craftdome/go-nym/contracts/mixnet/shared/models"
)

type GetPendingEpochEventsParams struct {
	StartAfter models.EpochEventID `json:"start_after,omitempty,omitzero"`
	Limit      uint32              `json:"limit,omitempty"`
}

type GetPendingIntervalEventsParams struct {
	StartAfter models.IntervalEventID `json:"start_after,omitempty,omitzero"`
	Limit      uint32                 `json:"limit,omitempty"`
}

type GetPendingEpochEventParams struct {
	EventID models.EpochEventID `json:"event_id"`
}

type GetPendingIntervalEventParams struct {
	EventID models.IntervalEventID `json:"event_id"`
}
