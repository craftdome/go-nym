package delegations

import (
	"context"
	"github.com/craftdome/go-nym/pkg/contracts/mixnet/models"
)

func (c *Client) GetByNode(ctx context.Context, params GetNodeDelegationsParams) (models.PagedNodeDelegations, error) {
	type req struct {
		MethodParams GetNodeDelegationsParams `json:"get_node_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	return Query[models.PagedNodeDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByDelegator(ctx context.Context, params GetDelegatorDelegationsParams) (models.PagedDelegatorDelegations, error) {
	type req struct {
		MethodParams GetDelegatorDelegationsParams `json:"get_delegator_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	return Query[models.PagedDelegatorDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByNodeAndDelegator(ctx context.Context, params GetNodeDelegationParams) (models.DelegatorNodeDelegation, error) {
	type req struct {
		MethodParams GetNodeDelegationParams `json:"get_delegation_details"`
	}

	r := req{MethodParams: params}

	return Query[models.DelegatorNodeDelegation](ctx, c.client, c.contract, r)
}

func (c *Client) GetAll(ctx context.Context, params GetDelegationsParams) (models.PagedAllDelegations, error) {
	type req struct {
		MethodParams GetDelegationsParams `json:"get_all_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, models.DelegationPageMaxRetrievalLimit)

	return Query[models.PagedAllDelegations](ctx, c.client, c.contract, r)
}
