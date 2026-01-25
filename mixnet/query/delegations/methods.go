package delegations

import (
	"context"

	"github.com/craftdome/go-nym/pkg/types"
)

func (c *Client) GetByNode(ctx context.Context, params GetNodeDelegationsParams) (types.PagedNodeDelegations, error) {
	type req struct {
		MethodParams GetNodeDelegationsParams `json:"get_node_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, types.DelegationPageMaxRetrievalLimit)

	return Query[types.PagedNodeDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByDelegator(ctx context.Context, params GetDelegatorDelegationsParams) (types.PagedDelegatorDelegations, error) {
	type req struct {
		MethodParams GetDelegatorDelegationsParams `json:"get_delegator_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, types.DelegationPageMaxRetrievalLimit)

	return Query[types.PagedDelegatorDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByNodeAndDelegator(ctx context.Context, params GetNodeDelegationParams) (types.DelegatorNodeDelegation, error) {
	type req struct {
		MethodParams GetNodeDelegationParams `json:"get_delegation_details"`
	}

	r := req{MethodParams: params}

	return Query[types.DelegatorNodeDelegation](ctx, c.client, c.contract, r)
}

func (c *Client) GetAll(ctx context.Context, params GetDelegationsParams) (types.PagedAllDelegations, error) {
	type req struct {
		MethodParams GetDelegationsParams `json:"get_all_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, types.DelegationPageMaxRetrievalLimit)

	return Query[types.PagedAllDelegations](ctx, c.client, c.contract, r)
}
