package mixnet

import (
	"github.com/craftdome/go-nym/mixnet/query/delegations"
	"github.com/craftdome/go-nym/mixnet/query/intervals"
	"github.com/craftdome/go-nym/mixnet/query/nodes"
	"github.com/craftdome/go-nym/mixnet/query/rewards"
)

// contract-related

// nodes-related

type GetAllBondedParams = nodes.GetAllBondedParams
type GetAllUnbondedParams = nodes.GetAllUnbondedParams
type GetUnbondedParams = nodes.GetUnbondedParams
type GetUnbondedByOwnerParams = nodes.GetUnbondedByOwnerParams
type GetUnbondedByIdentityKeyParams = nodes.GetUnbondedByIdentityKeyParams
type GetAllDetailedParams = nodes.GetAllDetailedParams
type GetDetailedParams = nodes.GetDetailedParams
type GetDetailedByOwnerParams = nodes.GetDetailedByOwnerParams
type GetDetailedByIdentityKeyParams = nodes.GetDetailedByIdentityKeyParams
type GetRewardingDetailsParams = nodes.GetRewardingDetailsParams
type GetStakeSaturationParams = nodes.GetStakeSaturationParams
type GetEpochAssignmentByRoleParams = nodes.GetEpochAssignmentByRoleParams

// delegations-related

type GetDelegationsParams = delegations.GetDelegationsParams
type GetNodeDelegationsParams = delegations.GetNodeDelegationsParams
type GetDelegatorDelegationsParams = delegations.GetDelegatorDelegationsParams
type GetNodeDelegationParams = delegations.GetNodeDelegationParams

// rewards-related

type GetPendingByOwnerParams = rewards.GetPendingByOwnerParams
type GetPendingByNodeParams = rewards.GetPendingByNodeParams
type GetPendingByNodeAndDelegatorParams = rewards.GetPendingByNodeAndDelegatorParams
type EstimateOperatorRewardParams = rewards.EstimateOperatorRewardParams
type EstimateDelegatorRewardParams = rewards.EstimateDelegatorRewardParams

// intervals-related

type GetPendingEpochEventsParams = intervals.GetPendingEpochEventsParams
type GetPendingEpochEventParams = intervals.GetPendingEpochEventParams
type GetPendingIntervalEventsParams = intervals.GetPendingIntervalEventsParams
type GetPendingIntervalEventParams = intervals.GetPendingIntervalEventParams
