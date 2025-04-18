package nodes

import (
	"github.com/craftdome/go-nym/pkg/types"
)

type GetAllBondedParams struct {
	StartAfter types.NodeID `json:"start_after,omitempty"`
	Limit      uint32       `json:"limit,omitempty"`
}

type GetAllDetailedParams struct {
	StartAfter types.NodeID `json:"start_after,omitempty"`
	Limit      uint32       `json:"limit,omitempty"`
}

type GetUnbondedParams struct {
	NodeID types.NodeID `json:"node_id"`
}

type GetAllUnbondedParams struct {
	StartAfter types.NodeID `json:"limit,omitempty"` // contract misspell: limit <-> start_after
	Limit      uint32       `json:"start_after,omitempty"`
}

type GetUnbondedByOwnerParams struct {
	Owner      types.Addr   `json:"owner"`
	StartAfter types.NodeID `json:"limit,omitempty"` // contract misspell: limit <-> start_after
	Limit      uint32       `json:"start_after,omitempty"`
}

type GetUnbondedByIdentityKeyParams struct {
	IdentityKey types.IdentityKey `json:"identity_key"`
	StartAfter  types.NodeID      `json:"start_after,omitempty"`
	Limit       uint32            `json:"limit,omitempty"`
}

type GetDetailedByOwnerParams struct {
	Owner types.Addr `json:"address"`
}

type GetDetailedParams struct {
	NodeID types.NodeID `json:"node_id"`
}

type GetDetailedByIdentityKeyParams struct {
	IdentityKey types.IdentityKey `json:"node_identity"`
}

type GetRewardingDetailsParams struct {
	NodeID types.NodeID `json:"node_id"`
}

type GetStakeSaturationParams struct {
	NodeID types.NodeID `json:"node_id"`
}

type GetEpochAssignmentByRoleParams struct {
	Role types.Role `json:"role"`
}
