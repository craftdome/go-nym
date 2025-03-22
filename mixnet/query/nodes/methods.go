package nodes

import (
	"context"
	"github.com/craftdome/go-nym/pkg/mixnet"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
)

func (c *Client) GetAllBonded(ctx context.Context, params GetAllBondedParams) (mixnet.PagedBondedNodes, error) {
	type req struct {
		MethodParams GetAllBondedParams `json:"get_nym_node_bonds_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.NymNodeBondMaxRetrievalLimit)

	return Query[mixnet.PagedBondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetAllDetailed(ctx context.Context, params GetAllDetailedParams) (mixnet.PagedDetailedNodes, error) {
	type req struct {
		MethodParams GetAllDetailedParams `json:"get_nym_nodes_detailed_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.NymNodeDetailsMaxRetrievalLimit)

	return Query[mixnet.PagedDetailedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetUnbonded(ctx context.Context, params GetUnbondedParams) (mixnet.UnbondedNode, error) {
	type req struct {
		MethodParams GetUnbondedParams `json:"get_unbonded_nym_node"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.UnbondedNode](ctx, c.client, c.contract, r)
}

// TODO start_after and limit dont work
func (c *Client) GetAllUnbonded(ctx context.Context, params GetAllUnbondedParams) (mixnet.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetAllUnbondedParams `json:"get_unbonded_nym_nodes_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.UnbondedNymNodesMaxRetrievalLimit)

	return Query[mixnet.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

// TODO start_next_after, limit do not work
func (c *Client) GetUnbondedByOwner(ctx context.Context, params GetUnbondedByOwnerParams) (mixnet.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByOwnerParams `json:"get_unbonded_nym_nodes_by_owner_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.UnbondedNymNodesMaxRetrievalLimit)

	return Query[mixnet.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetUnbondedByIdentityKey(ctx context.Context, params GetUnbondedByIdentityKeyParams) (mixnet.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByIdentityKeyParams `json:"get_unbonded_nym_nodes_by_identity_key_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, mixnet.UnbondedNymNodesMaxRetrievalLimit)

	return Query[mixnet.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetDetailedByOwner(ctx context.Context, params GetDetailedByOwnerParams) (mixnet.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByOwnerParams `json:"get_owned_nym_node"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		Operator mixnet.Addr          `json:"address"`
		Details  *mixnet.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return mixnet.DetailedNode{}, err
	}

	if resp.Details == nil {
		return mixnet.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetDetailed(ctx context.Context, params GetDetailedParams) (mixnet.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedParams `json:"get_nym_node_details"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		NodeID  mixnet.NodeID        `json:"node_id"`
		Details *mixnet.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return mixnet.DetailedNode{}, err
	}

	if resp.Details == nil {
		return mixnet.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetDetailedByIdentityKey(ctx context.Context, params GetDetailedByIdentityKeyParams) (mixnet.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByIdentityKeyParams `json:"get_nym_node_details_by_identity_key"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		IdentityKey mixnet.IdentityKey   `json:"identity_key"`
		Details     *mixnet.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return mixnet.DetailedNode{}, err
	}

	if resp.Details == nil {
		return mixnet.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetRewardingDetails(ctx context.Context, params GetRewardingDetailsParams) (mixnet.NodeRewardingDetails, error) {
	type req struct {
		MethodParams GetRewardingDetailsParams `json:"get_node_rewarding_details"`
	}

	r := req{MethodParams: params}

	type NodeRewardingDetailsResponse struct {
		NodeID           mixnet.NodeID               `json:"node_id"`
		RewardingDetails mixnet.NodeRewardingDetails `json:"rewarding_details"`
	}

	resp, err := Query[NodeRewardingDetailsResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return mixnet.NodeRewardingDetails{}, err
	}

	return resp.RewardingDetails, nil
}

func (c *Client) GetStakeSaturation(ctx context.Context, params GetStakeSaturationParams) (mixnet.NodeStakeSaturation, error) {
	type req struct {
		MethodParams GetStakeSaturationParams `json:"get_node_stake_saturation"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.NodeStakeSaturation](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochAssignmentByRole(ctx context.Context, params GetEpochAssignmentByRoleParams) (mixnet.EpochAssignment, error) {
	type req struct {
		MethodParams GetEpochAssignmentByRoleParams `json:"get_role_assignment"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.EpochAssignment](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochAssignmentMetadata(ctx context.Context) (mixnet.EpochAssignmentMetadata, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarded_set_metadata"`
	}

	r := req{}

	return Query[mixnet.EpochAssignmentMetadata](ctx, c.client, c.contract, r)
}
