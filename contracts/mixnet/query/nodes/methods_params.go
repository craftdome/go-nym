package nodes

import (
	"github.com/craftdome/go-nym/pkg/contracts/mixnet/models"
)

type GetAllBondedParams struct {
	StartAfter models.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetAllDetailedParams struct {
	StartAfter models.NodeID `json:"start_after"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedParams struct {
	NodeID models.NodeID `json:"node_id"`
}

type GetAllUnbondedParams struct {
	StartAfter models.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedByOwnerParams struct {
	Owner      models.Addr   `json:"owner"`
	StartAfter models.NodeID `json:"start_after,omitempty"`
	Limit      uint32        `json:"limit,omitempty"`
}

type GetUnbondedByIdentityKeyParams struct {
	IdentityKey models.IdentityKey `json:"identity_key"`
	StartAfter  models.NodeID      `json:"start_after,omitempty"`
	Limit       uint32             `json:"limit,omitempty"`
}

type GetDetailedByOwnerParams struct {
	Owner models.Addr `json:"address"`
}

type GetDetailedParams struct {
	NodeID models.NodeID `json:"node_id"`
}

type GetDetailedByIdentityKeyParams struct {
	IdentityKey models.IdentityKey `json:"node_identity"`
}

type GetRewardingDetailsParams struct {
	NodeID models.NodeID `json:"node_id"`
}

type GetStakeSaturationParams struct {
	NodeID models.NodeID `json:"node_id"`
}

type GetEpochAssignmentByRoleParams struct {
	Role models.Role `json:"role"`
}
