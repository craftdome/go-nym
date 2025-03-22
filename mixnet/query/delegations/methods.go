package delegations

import (
	"context"
	"github.com/craftdome/go-nym/pkg/mixnet"
)

func (c *Client) GetByNode(ctx context.Context, params GetNodeDelegationsParams) (mixnet.PagedNodeDelegations, error) {
	type req struct {
		MethodParams GetNodeDelegationsParams `json:"get_node_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, mixnet.DelegationPageMaxRetrievalLimit)

	return Query[mixnet.PagedNodeDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByDelegator(ctx context.Context, params GetDelegatorDelegationsParams) (mixnet.PagedDelegatorDelegations, error) {
	type req struct {
		MethodParams GetDelegatorDelegationsParams `json:"get_delegator_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, mixnet.DelegationPageMaxRetrievalLimit)

	return Query[mixnet.PagedDelegatorDelegations](ctx, c.client, c.contract, r)
}

func (c *Client) GetByNodeAndDelegator(ctx context.Context, params GetNodeDelegationParams) (mixnet.DelegatorNodeDelegation, error) {
	type req struct {
		MethodParams GetNodeDelegationParams `json:"get_delegation_details"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.DelegatorNodeDelegation](ctx, c.client, c.contract, r)
}

func (c *Client) GetAll(ctx context.Context, params GetDelegationsParams) (mixnet.PagedAllDelegations, error) {
	type req struct {
		MethodParams GetDelegationsParams `json:"get_all_delegations"`
	}

	r := req{MethodParams: params}
	r.MethodParams.Limit = min(r.MethodParams.Limit, mixnet.DelegationPageMaxRetrievalLimit)

	return Query[mixnet.PagedAllDelegations](ctx, c.client, c.contract, r)
}
