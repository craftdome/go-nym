package mixnet

import (
	"github.com/craftdome/go-nym/contracts/mixnet/query/delegations"
	"github.com/craftdome/go-nym/contracts/mixnet/query/intervals"
	"github.com/craftdome/go-nym/contracts/mixnet/query/nodes"
	"github.com/craftdome/go-nym/contracts/mixnet/query/rewards"
	"github.com/craftdome/go-nym/pkg/contracts/mixnet/models"
)

// contract-related

type ContractAdmin = models.ContractAdmin
type ContractVersion = models.ContractVersion
type ContractCW2Version = models.ContractCW2Version
type ContractStateParams = models.ContractStateParams
type ContractState = models.ContractState
type NodeVersion = models.NodeVersion
type PagedNodeVersionHistory = models.PagedNodeVersionHistory
type RewardingParams = models.RewardingParams
type EpochStatus = models.EpochStatus
type IntervalStatus = models.IntervalStatus

// nodes-related

type PagedBondedNodes = models.PagedBondedNodes
type PagedUnbondedNodes = models.PagedUnbondedNodes
type UnbondedNode = models.UnbondedNode
type PagedDetailedNodes = models.PagedDetailedNodes
type DetailedNode = models.DetailedNode
type NodeRewardingDetails = models.NodeRewardingDetails
type NodeStakeSaturation = models.NodeStakeSaturation
type EpochAssignment = models.EpochAssignment
type EpochAssignmentMetadata = models.EpochAssignmentMetadata

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

type PagedAllDelegations = models.PagedAllDelegations
type PagedNodeDelegations = models.PagedNodeDelegations
type PagedDelegatorDelegations = models.PagedDelegatorDelegations
type DelegatorNodeDelegation = models.DelegatorNodeDelegation

type GetDelegationsParams = delegations.GetDelegationsParams
type GetNodeDelegationsParams = delegations.GetNodeDelegationsParams
type GetDelegatorDelegationsParams = delegations.GetDelegatorDelegationsParams
type GetNodeDelegationParams = delegations.GetNodeDelegationParams

// rewards-related

type PendingReward = models.PendingReward
type EstimatedCurrentEpochReward = models.EstimatedCurrentEpochReward

type GetPendingByOwnerParams = rewards.GetPendingByOwnerParams
type GetPendingByNodeParams = rewards.GetPendingByNodeParams
type GetPendingByNodeAndDelegatorParams = rewards.GetPendingByNodeAndDelegatorParams
type EstimateOperatorRewardParams = rewards.EstimateOperatorRewardParams
type EstimateDelegatorRewardParams = rewards.EstimateDelegatorRewardParams

// intervals-related

type NumberOfPendingEvents = models.NumberOfPendingEvents
type PendingEpochEvents = models.PendingEpochEvents
type PendingEpochEvent = models.PendingEpochEvent
type PendingIntervalEvents = models.PendingIntervalEvents
type PendingIntervalEvent = models.PendingIntervalEvent

type GetPendingEpochEventsParams = intervals.GetPendingEpochEventsParams
type GetPendingEpochEventParams = intervals.GetPendingEpochEventParams
type GetPendingIntervalEventsParams = intervals.GetPendingIntervalEventsParams
type GetPendingIntervalEventParams = intervals.GetPendingIntervalEventParams
