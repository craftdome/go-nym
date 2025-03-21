package nodes

import (
	"context"
	"github.com/craftdome/go-nym/contracts/mixnet/models"
)

func (c *Client) GetAllBonded(ctx context.Context, params GetAllBondedParams) (*models.PagedBondedNodes, error) {
	type req struct {
		MethodParams GetAllBondedParams `json:"get_nym_node_bonds_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.NymNodeBondMaxRetrievalLimit)

	var result models.PagedBondedNodes
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetAllDetailed(ctx context.Context, params GetAllDetailedParams) (*models.PagedDetailedNodes, error) {
	type req struct {
		MethodParams GetAllDetailedParams `json:"get_nym_nodes_detailed_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.NymNodeDetailsMaxRetrievalLimit)

	var result models.PagedDetailedNodes
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetUnbonded(ctx context.Context, params GetUnbondedParams) (*models.UnbondedNode, error) {
	type req struct {
		MethodParams GetUnbondedParams `json:"get_unbonded_nym_node"`
	}

	r := req{MethodParams: params}

	var result models.UnbondedNode
	return &result, c.reader.Read(ctx, r, &result)
}

// TODO start_after and limit dont work
func (c *Client) GetAllUnbonded(ctx context.Context, params GetAllUnbondedParams) (*models.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetAllUnbondedParams `json:"get_unbonded_nym_nodes_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.UnbondedNymNodesMaxRetrievalLimit)

	var result models.PagedUnbondedNodes
	return &result, c.reader.Read(ctx, r, &result)
}

// TODO start_next_after, limit do not work
func (c *Client) GetUnbondedByOwner(ctx context.Context, params GetUnbondedByOwnerParams) (*models.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByOwnerParams `json:"get_unbonded_nym_nodes_by_owner_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.UnbondedNymNodesMaxRetrievalLimit)

	var result models.PagedUnbondedNodes
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetUnbondedByIdentityKey(ctx context.Context, params GetUnbondedByIdentityKeyParams) (*models.PagedUnbondedNodes, error) {
	type req struct {
		MethodParams GetUnbondedByIdentityKeyParams `json:"get_unbonded_nym_nodes_by_identity_key_paged"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(params.Limit, models.UnbondedNymNodesMaxRetrievalLimit)

	var result models.PagedUnbondedNodes
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetDetailedByOwner(ctx context.Context, params GetDetailedByOwnerParams) (*models.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByOwnerParams `json:"get_owned_nym_node"`
	}

	r := req{MethodParams: params}

	type resp struct {
		Operator models.Addr          `json:"address"`
		Details  *models.DetailedNode `json:"details"`
	}

	var result resp
	return result.Details, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetDetailed(ctx context.Context, params GetDetailedParams) (*models.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedParams `json:"get_nym_node_details"`
	}

	r := req{MethodParams: params}

	type resp struct {
		NodeID  models.NodeID        `json:"node_id"`
		Details *models.DetailedNode `json:"details"`
	}

	var result resp
	return result.Details, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetDetailedByIdentityKey(ctx context.Context, params GetDetailedByIdentityKeyParams) (*models.DetailedNode, error) {
	type req struct {
		MethodParams GetDetailedByIdentityKeyParams `json:"get_nym_node_details_by_identity_key"`
	}

	r := req{MethodParams: params}

	type resp struct {
		IdentityKey models.IdentityKey   `json:"identity_key"`
		Details     *models.DetailedNode `json:"details"`
	}

	var result resp
	return result.Details, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetRewardingDetails(ctx context.Context, params GetRewardingDetailsParams) (*models.NodeRewardingDetails, error) {
	type req struct {
		MethodParams GetRewardingDetailsParams `json:"get_node_rewarding_details"`
	}

	r := req{MethodParams: params}

	type resp struct {
		NodeID           models.NodeID               `json:"node_id"`
		RewardingDetails models.NodeRewardingDetails `json:"rewarding_details"`
	}

	var result resp
	return &result.RewardingDetails, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetStakeSaturation(ctx context.Context, params GetStakeSaturationParams) (*models.NodeStakeSaturation, error) {
	type req struct {
		MethodParams GetStakeSaturationParams `json:"get_node_stake_saturation"`
	}

	r := req{MethodParams: params}

	var result models.NodeStakeSaturation
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetEpochAssignmentByRole(ctx context.Context, params GetEpochAssignmentByRoleParams) (*models.EpochAssignment, error) {
	type req struct {
		MethodParams GetEpochAssignmentByRoleParams `json:"get_role_assignment"`
	}

	r := req{MethodParams: params}

	var result models.EpochAssignment
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetEpochAssignmentMetadata(ctx context.Context) (*models.EpochAssignmentMetadata, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarded_set_metadata"`
	}

	var result models.EpochAssignmentMetadata
	return &result, c.reader.Read(ctx, req{}, &result)
}
