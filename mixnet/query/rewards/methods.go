package rewards

import (
	"context"

	"github.com/craftdome/go-nym/pkg/types"
)

func (c *Client) GetPendingByOwner(ctx context.Context, params GetPendingByOwnerParams) (types.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByOwnerParams `json:"get_pending_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[types.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingByNode(ctx context.Context, params GetPendingByNodeParams) (types.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeParams `json:"get_pending_node_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[types.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingByNodeAndDelegator(ctx context.Context, params GetPendingByNodeAndDelegatorParams) (types.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeAndDelegatorParams `json:"get_pending_delegator_reward"`
	}

	r := req{MethodParams: params}

	return Query[types.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) EstimateOperatorReward(ctx context.Context, params EstimateOperatorRewardParams) (types.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateOperatorRewardParams `json:"get_estimated_current_epoch_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[types.EstimatedCurrentEpochReward](ctx, c.client, c.contract, r)
}

func (c *Client) EstimateDelegatorReward(ctx context.Context, params EstimateDelegatorRewardParams) (types.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateDelegatorRewardParams `json:"get_estimated_current_epoch_delegator_reward"`
	}

	r := req{MethodParams: params}

	return Query[types.EstimatedCurrentEpochReward](ctx, c.client, c.contract, r)
}
