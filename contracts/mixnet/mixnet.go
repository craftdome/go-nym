package mixnet

import (
	"context"
	"contracts/mixnet/query/contract"
	"contracts/mixnet/query/delegations"
	"contracts/mixnet/query/intervals"
	"contracts/mixnet/query/nodes"
	"contracts/mixnet/query/rewards"
	"contracts/mixnet/shared/models"
	"github.com/CosmWasm/wasmd/x/wasm/types"
	"google.golang.org/grpc"
)

type Contract interface {
	GetAdmin(context.Context) (*ContractAdmin, error)
	GetVersion(context.Context) (*ContractVersion, error)
	GetCW2Version(context.Context) (*ContractCW2Version, error)
	GetStateParams(context.Context) (*ContractStateParams, error)
	GetState(context.Context) (*ContractState, error)
	GetCurrentNodeVersion(context.Context) (*NodeVersion, error)
	GetNodeVersionHistory(context.Context) (*PagedNodeVersionHistory, error)
	GetRewardingParams(context.Context) (*RewardingParams, error)
	GetEpochStatus(context.Context) (*EpochStatus, error)
	GetIntervalStatus(context.Context) (*IntervalStatus, error)
}

type Nodes interface {
	GetAllBonded(context.Context, nodes.GetAllBondedParams) (*PagedBondedNodes, error)

	GetAllUnbonded(context.Context, nodes.GetAllUnbondedParams) (*PagedUnbondedNodes, error)
	GetUnbonded(context.Context, nodes.GetUnbondedParams) (*UnbondedNode, error)
	GetUnbondedByOwner(context.Context, nodes.GetUnbondedByOwnerParams) (*PagedUnbondedNodes, error)
	GetUnbondedByIdentityKey(context.Context, nodes.GetUnbondedByIdentityKeyParams) (*PagedUnbondedNodes, error)

	GetAllDetailed(context.Context, nodes.GetAllDetailedParams) (*PagedDetailedNodes, error)
	GetDetailed(context.Context, nodes.GetDetailedParams) (*DetailedNode, error)
	GetDetailedByOwner(context.Context, nodes.GetDetailedByOwnerParams) (*DetailedNode, error)
	GetDetailedByIdentityKey(context.Context, nodes.GetDetailedByIdentityKeyParams) (*DetailedNode, error)

	GetRewardingDetails(context.Context, nodes.GetRewardingDetailsParams) (*NodeRewardingDetails, error)
	GetStakeSaturation(context.Context, nodes.GetStakeSaturationParams) (*NodeStakeSaturation, error)

	GetEpochAssignmentByRole(context.Context, nodes.GetEpochAssignmentByRoleParams) (*EpochAssignment, error)
	GetEpochAssignmentMetadata(context.Context) (*EpochAssignmentMetadata, error)
}

type Delegations interface {
	GetAll(context.Context, delegations.GetDelegationsParams) (*PagedAllDelegations, error)
	GetByNode(context.Context, delegations.GetNodeDelegationsParams) (*PagedNodeDelegations, error)
	GetByDelegator(context.Context, delegations.GetDelegatorDelegationsParams) (*PagedDelegatorDelegations, error)
	GetByNodeAndDelegator(context.Context, delegations.GetNodeDelegationParams) (*DelegatorNodeDelegation, error)
}

type Rewards interface {
	GetPendingByOwner(context.Context, rewards.GetPendingByOwnerParams) (*PendingReward, error)
	GetPendingByNode(context.Context, rewards.GetPendingByNodeParams) (*PendingReward, error)
	GetPendingByNodeAndDelegator(context.Context, rewards.GetPendingByNodeAndDelegatorParams) (*PendingReward, error)
	EstimateOperatorReward(context.Context, rewards.EstimateOperatorRewardParams) (*EstimatedCurrentEpochReward, error)
	EstimateDelegatorReward(context.Context, rewards.EstimateDelegatorRewardParams) (*EstimatedCurrentEpochReward, error)
}

type Intervals interface {
	GetNumberOfPendingEvents(context.Context) (*models.NumberOfPendingEvents, error)
	GetPendingEpochEvents(context.Context, intervals.GetPendingEpochEventsParams) (*PendingEpochEvents, error)
	GetPendingEpochEvent(context.Context, intervals.GetPendingEpochEventParams) (*PendingEpochEvent, error)
	GetPendingIntervalEvents(context.Context, intervals.GetPendingIntervalEventsParams) (*PendingIntervalEvents, error)
	GetPendingIntervalEvent(context.Context, intervals.GetPendingIntervalEventParams) (*PendingIntervalEvent, error)
}

type QueryClient struct {
	Contract    Contract
	Nodes       Nodes
	Rewards     Rewards
	Intervals   Intervals
	Delegations Delegations
}

func NewQueryClient(conn *grpc.ClientConn, contractAddress string) *QueryClient {
	queryClient := &client{
		client:   types.NewQueryClient(conn),
		contract: contractAddress,
	}

	return &QueryClient{
		Contract:    contract.New(queryClient),
		Nodes:       nodes.New(queryClient),
		Rewards:     rewards.New(queryClient),
		Intervals:   intervals.New(queryClient),
		Delegations: delegations.New(queryClient),
	}
}
