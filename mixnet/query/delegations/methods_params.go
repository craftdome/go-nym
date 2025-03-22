package delegations

import (
	"github.com/craftdome/go-nym/pkg/mixnet"
)

type GetNodeDelegationsParams struct {
	NodeID     mixnet.NodeID `json:"node_id"`
	StartAfter mixnet.Addr   `json:"start_after,omitzero"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetDelegatorDelegationsParams struct {
	Delegator  mixnet.Addr        `json:"delegator"`
	StartAfter *mixnet.StorageKey `json:"start_after"`
	Limit      uint32             `json:"limit,omitempty"`
}

type GetNodeDelegationParams struct {
	NodeID    mixnet.NodeID `json:"node_id"`
	Delegator mixnet.Addr   `json:"delegator"`
	Proxy     string        `json:"proxy,omitempty"`
}

type GetDelegationsParams struct {
	StartAfter *mixnet.StorageKey `json:"start_after"`
	Limit      uint32             `json:"limit,omitempty"`
}
