package rewards

import (
	"context"
	"github.com/craftdome/go-nym/pkg/mixnet"
)

func (c *Client) GetPendingByOwner(ctx context.Context, params GetPendingByOwnerParams) (mixnet.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByOwnerParams `json:"get_pending_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingByNode(ctx context.Context, params GetPendingByNodeParams) (mixnet.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeParams `json:"get_pending_node_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) GetPendingByNodeAndDelegator(ctx context.Context, params GetPendingByNodeAndDelegatorParams) (mixnet.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeAndDelegatorParams `json:"get_pending_delegator_reward"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.PendingReward](ctx, c.client, c.contract, r)
}

func (c *Client) EstimateOperatorReward(ctx context.Context, params EstimateOperatorRewardParams) (mixnet.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateOperatorRewardParams `json:"get_estimated_current_epoch_operator_reward"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.EstimatedCurrentEpochReward](ctx, c.client, c.contract, r)
}

func (c *Client) EstimateDelegatorReward(ctx context.Context, params EstimateDelegatorRewardParams) (mixnet.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateDelegatorRewardParams `json:"get_estimated_current_epoch_delegator_reward"`
	}

	r := req{MethodParams: params}

	return Query[mixnet.EstimatedCurrentEpochReward](ctx, c.client, c.contract, r)
}
