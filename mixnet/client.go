package mixnet

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"google.golang.org/grpc"

	"github.com/craftdome/go-nym/mixnet/query/contract"
	"github.com/craftdome/go-nym/mixnet/query/delegations"
	"github.com/craftdome/go-nym/mixnet/query/intervals"
	"github.com/craftdome/go-nym/mixnet/query/nodes"
	"github.com/craftdome/go-nym/mixnet/query/rewards"
	"github.com/craftdome/go-nym/pkg/types"
)

type Contract interface {
	GetAdmin(context.Context) (types.ContractAdmin, error)
	GetVersion(context.Context) (types.ContractVersion, error)
	GetCW2Version(context.Context) (types.ContractCW2Version, error)
	GetStateParams(context.Context) (types.ContractStateParams, error)
	GetState(context.Context) (types.ContractState, error)
	GetCurrentNodeVersion(context.Context) (types.NodeVersion, error)
	GetNodeVersionHistory(context.Context) (types.PagedNodeVersionHistory, error)
	GetRewardingParams(context.Context) (types.RewardingParams, error)
	GetEpochStatus(context.Context) (types.EpochStatus, error)
	GetIntervalStatus(context.Context) (types.IntervalStatus, error)
}

type Nodes interface {
	GetAllBonded(context.Context, GetAllBondedParams) (types.PagedBondedNodes, error)

	GetAllUnbonded(context.Context, GetAllUnbondedParams) (types.PagedUnbondedNodes, error)
	GetUnbonded(context.Context, GetUnbondedParams) (types.UnbondedNode, error)
	GetUnbondedByOwner(context.Context, GetUnbondedByOwnerParams) (types.PagedUnbondedNodes, error)
	GetUnbondedByIdentityKey(context.Context, GetUnbondedByIdentityKeyParams) (types.PagedUnbondedNodes, error)

	GetAllDetailed(context.Context, GetAllDetailedParams) (types.PagedDetailedNodes, error)
	GetDetailed(context.Context, GetDetailedParams) (types.DetailedNode, error)
	GetDetailedByOwner(context.Context, GetDetailedByOwnerParams) (types.DetailedNode, error)
	GetDetailedByIdentityKey(context.Context, GetDetailedByIdentityKeyParams) (types.DetailedNode, error)

	GetRewardingDetails(context.Context, GetRewardingDetailsParams) (types.NodeRewardingDetails, error)
	GetStakeSaturation(context.Context, GetStakeSaturationParams) (types.NodeStakeSaturation, error)

	GetEpochAssignmentByRole(context.Context, GetEpochAssignmentByRoleParams) (types.EpochAssignment, error)
	GetEpochAssignmentMetadata(context.Context) (types.EpochAssignmentMetadata, error)
}

type Delegations interface {
	GetAll(context.Context, GetDelegationsParams) (types.PagedAllDelegations, error)
	GetByNode(context.Context, GetNodeDelegationsParams) (types.PagedNodeDelegations, error)
	GetByDelegator(context.Context, GetDelegatorDelegationsParams) (types.PagedDelegatorDelegations, error)
	GetByNodeAndDelegator(context.Context, GetNodeDelegationParams) (types.DelegatorNodeDelegation, error)
}

type Rewards interface {
	GetPendingByOwner(context.Context, GetPendingByOwnerParams) (types.PendingReward, error)
	GetPendingByNode(context.Context, GetPendingByNodeParams) (types.PendingReward, error)
	GetPendingByNodeAndDelegator(context.Context, GetPendingByNodeAndDelegatorParams) (types.PendingReward, error)
	EstimateOperatorReward(context.Context, EstimateOperatorRewardParams) (types.EstimatedCurrentEpochReward, error)
	EstimateDelegatorReward(context.Context, EstimateDelegatorRewardParams) (types.EstimatedCurrentEpochReward, error)
}

type Intervals interface {
	GetNumberOfPendingEvents(context.Context) (types.NumberOfPendingEvents, error)
	GetPendingEpochEvents(context.Context, GetPendingEpochEventsParams) (types.PendingEpochEvents, error)
	GetPendingEpochEvent(context.Context, GetPendingEpochEventParams) (types.PendingEpochEvent, error)
	GetPendingIntervalEvents(context.Context, GetPendingIntervalEventsParams) (types.PendingIntervalEvents, error)
	GetPendingIntervalEvent(context.Context, GetPendingIntervalEventParams) (types.PendingIntervalEvent, error)
}

type Client struct {
	Contract    Contract
	Nodes       Nodes
	Rewards     Rewards
	Intervals   Intervals
	Delegations Delegations
}

func New(conn *grpc.ClientConn, contractAddress string) *Client {
	queryClient := wasmtypes.NewQueryClient(conn)

	return &Client{
		Contract:    contract.New(queryClient, contractAddress),
		Nodes:       nodes.New(queryClient, contractAddress),
		Rewards:     rewards.New(queryClient, contractAddress),
		Intervals:   intervals.New(queryClient, contractAddress),
		Delegations: delegations.New(queryClient, contractAddress),
	}
}
