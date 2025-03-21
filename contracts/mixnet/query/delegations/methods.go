package delegations

import (
	"context"
	"github.com/craftdome/go-nym/contracts/mixnet/models"
)

func (c *Client) GetByNode(ctx context.Context, params GetNodeDelegationsParams) (*models.PagedNodeDelegations, error) {
	type req struct {
		MethodParams GetNodeDelegationsParams `json:"get_node_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	var result models.PagedNodeDelegations
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetByDelegator(ctx context.Context, params GetDelegatorDelegationsParams) (*models.PagedDelegatorDelegations, error) {
	type req struct {
		MethodParams GetDelegatorDelegationsParams `json:"get_delegator_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	var result models.PagedDelegatorDelegations
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetByNodeAndDelegator(ctx context.Context, params GetNodeDelegationParams) (*models.DelegatorNodeDelegation, error) {
	type req struct {
		MethodParams GetNodeDelegationParams `json:"get_delegation_details"`
	}

	r := req{MethodParams: params}

	var result models.DelegatorNodeDelegation
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetAll(ctx context.Context, params GetDelegationsParams) (*models.PagedAllDelegations, error) {
	type req struct {
		MethodParams GetDelegationsParams `json:"get_all_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	var result models.PagedAllDelegations
	return &result, c.reader.Read(ctx, r, &result)
}
