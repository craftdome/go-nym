package mixnet

import (
	"context"
	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/craftdome/go-nym/mixnet/query/contract"
	"github.com/craftdome/go-nym/mixnet/query/delegations"
	"github.com/craftdome/go-nym/mixnet/query/intervals"
	"github.com/craftdome/go-nym/mixnet/query/nodes"
	"github.com/craftdome/go-nym/mixnet/query/rewards"
	"google.golang.org/grpc"
)

type Contract interface {
	GetAdmin(context.Context) (ContractAdmin, error)
	GetVersion(context.Context) (ContractVersion, error)
	GetCW2Version(context.Context) (ContractCW2Version, error)
	GetStateParams(context.Context) (ContractStateParams, error)
	GetState(context.Context) (ContractState, error)
	GetCurrentNodeVersion(context.Context) (NodeVersion, error)
	GetNodeVersionHistory(context.Context) (PagedNodeVersionHistory, error)
	GetRewardingParams(context.Context) (RewardingParams, error)
	GetEpochStatus(context.Context) (EpochStatus, error)
	GetIntervalStatus(context.Context) (IntervalStatus, error)
}

type Nodes interface {
	GetAllBonded(context.Context, GetAllBondedParams) (PagedBondedNodes, error)

	GetAllUnbonded(context.Context, GetAllUnbondedParams) (PagedUnbondedNodes, error)
	GetUnbonded(context.Context, GetUnbondedParams) (UnbondedNode, error)
	GetUnbondedByOwner(context.Context, GetUnbondedByOwnerParams) (PagedUnbondedNodes, error)
	GetUnbondedByIdentityKey(context.Context, GetUnbondedByIdentityKeyParams) (PagedUnbondedNodes, error)

	GetAllDetailed(context.Context, GetAllDetailedParams) (PagedDetailedNodes, error)
	GetDetailed(context.Context, GetDetailedParams) (DetailedNode, error)
	GetDetailedByOwner(context.Context, GetDetailedByOwnerParams) (DetailedNode, error)
	GetDetailedByIdentityKey(context.Context, GetDetailedByIdentityKeyParams) (DetailedNode, error)

	GetRewardingDetails(context.Context, GetRewardingDetailsParams) (NodeRewardingDetails, error)
	GetStakeSaturation(context.Context, GetStakeSaturationParams) (NodeStakeSaturation, error)

	GetEpochAssignmentByRole(context.Context, GetEpochAssignmentByRoleParams) (EpochAssignment, error)
	GetEpochAssignmentMetadata(context.Context) (EpochAssignmentMetadata, error)
}

type Delegations interface {
	GetAll(context.Context, GetDelegationsParams) (PagedAllDelegations, error)
	GetByNode(context.Context, GetNodeDelegationsParams) (PagedNodeDelegations, error)
	GetByDelegator(context.Context, GetDelegatorDelegationsParams) (PagedDelegatorDelegations, error)
	GetByNodeAndDelegator(context.Context, GetNodeDelegationParams) (DelegatorNodeDelegation, error)
}

type Rewards interface {
	GetPendingByOwner(context.Context, GetPendingByOwnerParams) (PendingReward, error)
	GetPendingByNode(context.Context, GetPendingByNodeParams) (PendingReward, error)
	GetPendingByNodeAndDelegator(context.Context, GetPendingByNodeAndDelegatorParams) (PendingReward, error)
	EstimateOperatorReward(context.Context, EstimateOperatorRewardParams) (EstimatedCurrentEpochReward, error)
	EstimateDelegatorReward(context.Context, EstimateDelegatorRewardParams) (EstimatedCurrentEpochReward, error)
}

type Intervals interface {
	GetNumberOfPendingEvents(context.Context) (NumberOfPendingEvents, error)
	GetPendingEpochEvents(context.Context, GetPendingEpochEventsParams) (PendingEpochEvents, error)
	GetPendingEpochEvent(context.Context, GetPendingEpochEventParams) (PendingEpochEvent, error)
	GetPendingIntervalEvents(context.Context, GetPendingIntervalEventsParams) (PendingIntervalEvents, error)
	GetPendingIntervalEvent(context.Context, GetPendingIntervalEventParams) (PendingIntervalEvent, error)
}

type Client struct {
	Contract    Contract
	Nodes       Nodes
	Rewards     Rewards
	Intervals   Intervals
	Delegations Delegations
}

func New(conn *grpc.ClientConn, contractAddress string) *Client {
	queryClient := types.NewQueryClient(conn)

	return &Client{
		Contract:    contract.New(queryClient, contractAddress),
		Nodes:       nodes.New(queryClient, contractAddress),
		Rewards:     rewards.New(queryClient, contractAddress),
		Intervals:   intervals.New(queryClient, contractAddress),
		Delegations: delegations.New(queryClient, contractAddress),
	}
}
