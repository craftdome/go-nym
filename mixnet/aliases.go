package mixnet

import (
	"github.com/craftdome/go-nym/mixnet/query/delegations"
	"github.com/craftdome/go-nym/mixnet/query/intervals"
	"github.com/craftdome/go-nym/mixnet/query/nodes"
	"github.com/craftdome/go-nym/mixnet/query/rewards"
	"github.com/craftdome/go-nym/pkg/mixnet"
)

var (
	TokenSupply                        = mixnet.TokenSupply
	DefaultIntervalOperatingCostAmount = mixnet.DefaultIntervalOperatingCostAmount
	DefaultProfitMarginPercent         = mixnet.DefaultProfitMarginPercent
	UnitDelegationBase                 = mixnet.UnitDelegationBase
)

const (
	NymNodeBondDefaultRetrievalLimit = mixnet.NymNodeBondDefaultRetrievalLimit
	NymNodeBondMaxRetrievalLimit     = mixnet.NymNodeBondMaxRetrievalLimit

	NymNodeDetailsDefaultRetrievalLimit = mixnet.NymNodeDetailsDefaultRetrievalLimit
	NymNodeDetailsMaxRetrievalLimit     = mixnet.NymNodeDetailsMaxRetrievalLimit

	UnbondedNymNodesDefaultRetrievalLimit = mixnet.UnbondedNymNodesDefaultRetrievalLimit
	UnbondedNymNodesMaxRetrievalLimit     = mixnet.UnbondedNymNodesMaxRetrievalLimit

	DelegationPageDefaultRetrievalLimit = mixnet.DelegationPageDefaultRetrievalLimit
	DelegationPageMaxRetrievalLimit     = mixnet.DelegationPageMaxRetrievalLimit

	EpochEventsDefaultRetrievalLimit = mixnet.EpochEventsDefaultRetrievalLimit
	EpochEventsMaxRetrievalLimit     = mixnet.EpochEventsMaxRetrievalLimit

	IntervalEventsDefaultRetrievalLimit = mixnet.IntervalEventsDefaultRetrievalLimit
	IntervalEventsMaxRetrievalLimit     = mixnet.IntervalEventsMaxRetrievalLimit

	EntryGatewayRole = mixnet.EntryGateway
	Layer1Role       = mixnet.Layer1
	Layer2Role       = mixnet.Layer2
	Layer3Role       = mixnet.Layer3
	ExitGatewayRole  = mixnet.ExitGateway
	StandbyRole      = mixnet.Standby
)

type Coin = mixnet.Coin
type Percent = mixnet.Percent
type Decimal = mixnet.Decimal
type Addr = mixnet.Addr
type NodeID = mixnet.NodeID
type Role = mixnet.Role
type IntervalID = mixnet.IntervalID
type EpochID = mixnet.EpochID
type IdentityKey = mixnet.IdentityKey
type EpochEventID = mixnet.EpochEventID
type IntervalEventID = mixnet.IntervalEventID
type BlockHeight = mixnet.BlockHeight

// contract-related

type ContractAdmin = mixnet.ContractAdmin
type ContractVersion = mixnet.ContractVersion
type ContractCW2Version = mixnet.ContractCW2Version
type ContractStateParams = mixnet.ContractStateParams
type ProfitMarginRange = mixnet.ProfitMarginRange
type OperatingCostRange = mixnet.OperatingCostRange
type ContractState = mixnet.ContractState
type NodeVersion = mixnet.NodeVersion
type NodeVersionInfo = mixnet.NodeVersionInfo
type PagedNodeVersionHistory = mixnet.PagedNodeVersionHistory
type RewardingParams = mixnet.RewardingParams
type RewardedSetParams = mixnet.RewardedSetParams
type EpochStatus = mixnet.EpochStatus
type InProgressEpochState = mixnet.InProgressEpochState
type RewardingEpochState = mixnet.RewardingEpochState
type ReconcilingEventsEpochState = mixnet.ReconcilingEventsEpochState
type RoleAssignmentEpochState = mixnet.RoleAssignmentEpochState
type IntervalStatus = mixnet.IntervalStatus
type Interval = mixnet.Interval
type OffsetDateTime = mixnet.OffsetDateTime

// nodes-related

type PagedBondedNodes = mixnet.PagedBondedNodes
type BondedNode = mixnet.BondedNode
type Node = mixnet.Node
type PagedUnbondedNodes = mixnet.PagedUnbondedNodes
type UnbondedNode = mixnet.UnbondedNode
type PagedDetailedNodes = mixnet.PagedDetailedNodes
type DetailedNode = mixnet.DetailedNode
type NodeRewardingDetails = mixnet.NodeRewardingDetails
type NodeCostParams = mixnet.NodeCostParams
type NodeStakeSaturation = mixnet.NodeStakeSaturation
type EpochAssignment = mixnet.EpochAssignment
type EpochAssignmentMetadata = mixnet.EpochAssignmentMetadata
type RewardedSetMetadata = mixnet.RewardedSetMetadata
type RoleMetadata = mixnet.RoleMetadata

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

type PagedAllDelegations = mixnet.PagedAllDelegations
type StorageKey = mixnet.StorageKey
type Delegation = mixnet.Delegation
type PagedNodeDelegations = mixnet.PagedNodeDelegations
type OwnerProxySubKey = mixnet.OwnerProxySubKey
type PagedDelegatorDelegations = mixnet.PagedDelegatorDelegations
type DelegatorNodeDelegation = mixnet.DelegatorNodeDelegation

type GetDelegationsParams = delegations.GetDelegationsParams
type GetNodeDelegationsParams = delegations.GetNodeDelegationsParams
type GetDelegatorDelegationsParams = delegations.GetDelegatorDelegationsParams
type GetNodeDelegationParams = delegations.GetNodeDelegationParams

// rewards-related

type PendingReward = mixnet.PendingReward
type EstimatedCurrentEpochReward = mixnet.EstimatedCurrentEpochReward

type GetPendingByOwnerParams = rewards.GetPendingByOwnerParams
type GetPendingByNodeParams = rewards.GetPendingByNodeParams
type GetPendingByNodeAndDelegatorParams = rewards.GetPendingByNodeAndDelegatorParams
type EstimateOperatorRewardParams = rewards.EstimateOperatorRewardParams
type EstimateDelegatorRewardParams = rewards.EstimateDelegatorRewardParams

// intervals-related

type NumberOfPendingEvents = mixnet.NumberOfPendingEvents
type PendingEpochEvents = mixnet.PendingEpochEvents
type PendingEpochEventData = mixnet.PendingEpochEventData
type PendingEpochEventKind = mixnet.PendingEpochEventKind
type ActiveSetUpdate = mixnet.ActiveSetUpdate
type PendingEpochEvent = mixnet.PendingEpochEvent
type PendingIntervalEvents = mixnet.PendingIntervalEvents
type PendingIntervalEventData = mixnet.PendingIntervalEventData
type IntervalRewardingParamsUpdate = mixnet.IntervalRewardingParamsUpdate
type PendingIntervalEvent = mixnet.PendingIntervalEvent

type GetPendingEpochEventsParams = intervals.GetPendingEpochEventsParams
type GetPendingEpochEventParams = intervals.GetPendingEpochEventParams
type GetPendingIntervalEventsParams = intervals.GetPendingIntervalEventsParams
type GetPendingIntervalEventParams = intervals.GetPendingIntervalEventParams
