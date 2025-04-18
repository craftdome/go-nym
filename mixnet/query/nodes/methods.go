package nodes

import (
	"context"
	"github.com/craftdome/go-nym/pkg/types"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
)

func (c *Client) GetAllBonded(ctx context.Context, params GetAllBondedParams) (types.PagedBondedNodes, error) {
	type req struct {
		MethodParams GetAllBondedParams `json:"get_nym_node_bonds_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.NymNodeBondMaxRetrievalLimit)

	return Query[types.PagedBondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetAllDetailed(ctx context.Context, params GetAllDetailedParams) (types.PagedDetailedNodes, error) {
	type req struct {
		MethodParams GetAllDetailedParams `json:"get_nym_nodes_detailed_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.NymNodeDetailsMaxRetrievalLimit)

	return Query[types.PagedDetailedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetUnbonded(ctx context.Context, params GetUnbondedParams) (types.UnbondedNode, error) {
	type req struct {
		MethodParams GetUnbondedParams `json:"get_unbonded_nym_node"`
	}

	r := req{MethodParams: params}

	type UnbondedNodeResponse struct {
		NodeID  types.NodeID        `json:"node_id"`
		Details *types.UnbondedNode `json:"details"`
	}

	resp, err := Query[UnbondedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return types.UnbondedNode{}, err
	}

	if resp.Details == nil {
		return types.UnbondedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetAllUnbonded(ctx context.Context, params GetAllUnbondedParams) (types.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetAllUnbondedParams `json:"get_unbonded_nym_nodes_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.UnbondedNymNodesMaxRetrievalLimit)

	return Query[types.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetUnbondedByOwner(ctx context.Context, params GetUnbondedByOwnerParams) (types.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByOwnerParams `json:"get_unbonded_nym_nodes_by_owner_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.UnbondedNymNodesMaxRetrievalLimit)

	return Query[types.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetUnbondedByIdentityKey(ctx context.Context, params GetUnbondedByIdentityKeyParams) (types.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByIdentityKeyParams `json:"get_unbonded_nym_nodes_by_identity_key_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, types.UnbondedNymNodesMaxRetrievalLimit)

	return Query[types.PagedUnbondedNodes](ctx, c.client, c.contract, r)
}

func (c *Client) GetDetailedByOwner(ctx context.Context, params GetDetailedByOwnerParams) (types.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByOwnerParams `json:"get_owned_nym_node"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		Operator types.Addr          `json:"address"`
		Details  *types.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return types.DetailedNode{}, err
	}

	if resp.Details == nil {
		return types.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetDetailed(ctx context.Context, params GetDetailedParams) (types.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedParams `json:"get_nym_node_details"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		NodeID  types.NodeID        `json:"node_id"`
		Details *types.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return types.DetailedNode{}, err
	}

	if resp.Details == nil {
		return types.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetDetailedByIdentityKey(ctx context.Context, params GetDetailedByIdentityKeyParams) (types.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByIdentityKeyParams `json:"get_nym_node_details_by_identity_key"`
	}

	r := req{MethodParams: params}

	type DetailedNodeResponse struct {
		IdentityKey types.IdentityKey   `json:"identity_key"`
		Details     *types.DetailedNode `json:"details"`
	}

	resp, err := Query[DetailedNodeResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return types.DetailedNode{}, err
	}

	if resp.Details == nil {
		return types.DetailedNode{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetRewardingDetails(ctx context.Context, params GetRewardingDetailsParams) (types.NodeRewardingDetails, error) {
	type req struct {
		MethodParams GetRewardingDetailsParams `json:"get_node_rewarding_details"`
	}

	r := req{MethodParams: params}

	type NodeRewardingDetailsResponse struct {
		NodeID  types.NodeID                `json:"node_id"`
		Details *types.NodeRewardingDetails `json:"rewarding_details"`
	}

	resp, err := Query[NodeRewardingDetailsResponse](ctx, c.client, c.contract, r)
	if err != nil {
		return types.NodeRewardingDetails{}, err
	}

	if resp.Details == nil {
		return types.NodeRewardingDetails{}, ErrNotFound
	}

	return *resp.Details, nil
}

func (c *Client) GetStakeSaturation(ctx context.Context, params GetStakeSaturationParams) (types.NodeStakeSaturation, error) {
	type req struct {
		MethodParams GetStakeSaturationParams `json:"get_node_stake_saturation"`
	}

	r := req{MethodParams: params}

	return Query[types.NodeStakeSaturation](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochAssignmentByRole(ctx context.Context, params GetEpochAssignmentByRoleParams) (types.EpochAssignment, error) {
	type req struct {
		MethodParams GetEpochAssignmentByRoleParams `json:"get_role_assignment"`
	}

	r := req{MethodParams: params}

	return Query[types.EpochAssignment](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochAssignmentMetadata(ctx context.Context) (types.EpochAssignmentMetadata, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarded_set_metadata"`
	}

	r := req{}

	return Query[types.EpochAssignmentMetadata](ctx, c.client, c.contract, r)
}
