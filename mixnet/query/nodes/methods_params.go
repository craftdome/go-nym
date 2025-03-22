package nodes

import (
	"github.com/craftdome/go-nym/pkg/mixnet"
)

type GetAllBondedParams struct {
	StartAfter mixnet.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetAllDetailedParams struct {
	StartAfter mixnet.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedParams struct {
	NodeID mixnet.NodeID `json:"node_id"`
}

type GetAllUnbondedParams struct {
	StartAfter mixnet.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedByOwnerParams struct {
	Owner      mixnet.Addr   `json:"owner"`
	StartAfter mixnet.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedByIdentityKeyParams struct {
	IdentityKey mixnet.IdentityKey `json:"identity_key"`
	StartAfter  mixnet.NodeID      `json:"start_after,omitempty"`
	Limit       uint32             `json:"limit,omitempty"`
}

type GetDetailedByOwnerParams struct {
	Owner mixnet.Addr `json:"address"`
}

type GetDetailedParams struct {
	NodeID mixnet.NodeID `json:"node_id"`
}

type GetDetailedByIdentityKeyParams struct {
	IdentityKey mixnet.IdentityKey `json:"node_identity"`
}

type GetRewardingDetailsParams struct {
	NodeID mixnet.NodeID `json:"node_id"`
}

type GetStakeSaturationParams struct {
	NodeID mixnet.NodeID `json:"node_id"`
}

type GetEpochAssignmentByRoleParams struct {
	Role mixnet.Role `json:"role"`
}
