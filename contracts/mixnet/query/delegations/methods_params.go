package delegations

import (
	"contracts/mixnet/shared/models"
)

type GetNodeDelegationsParams struct {
	NodeID     models.NodeID `json:"node_id"`
	StartAfter models.Addr   `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetDelegatorDelegationsParams struct {
	Delegator  models.Addr       `json:"delegator"`
	StartAfter models.StorageKey `json:"start_after,omitempty"`
	Limit      uint32            `json:"limit,omitempty"`
}

type GetNodeDelegationParams struct {
	NodeID    models.NodeID `json:"node_id"`
	Delegator models.Addr   `json:"delegator"`
	Proxy     string        `json:"proxy,omitempty"`
}

type GetDelegationsParams struct {
	StartAfter models.StorageKey `json:"start_after,omitempty"`
	Limit      uint32            `json:"limit,omitempty"`
}
