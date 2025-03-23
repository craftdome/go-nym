package delegations

import (
	"github.com/craftdome/go-nym/pkg/types"
)

type GetNodeDelegationsParams struct {
	NodeID     types.NodeID `json:"node_id"`
	StartAfter types.Addr   `json:"start_after,omitzero"`
	Limit      uint32       `json:"limit,omitempty"`
}

type GetDelegatorDelegationsParams struct {
	Delegator  types.Addr        `json:"delegator"`
	StartAfter *types.StorageKey `json:"start_after"`
	Limit      uint32            `json:"limit,omitempty"`
}

type GetNodeDelegationParams struct {
	NodeID    types.NodeID `json:"node_id"`
	Delegator types.Addr   `json:"delegator"`
	Proxy     string       `json:"proxy,omitempty"`
}

type GetDelegationsParams struct {
	StartAfter *types.StorageKey `json:"start_after"`
	Limit      uint32            `json:"limit,omitempty"`
}
