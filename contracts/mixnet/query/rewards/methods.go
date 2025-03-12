package rewards

import (
	"context"
	"github.com/craftdome/go-nym/contracts/v1/mixnet/shared/models"
)

func (c *Client) GetPendingByOwner(ctx context.Context, params GetPendingByOwnerParams) (*models.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByOwnerParams `json:"get_pending_operator_reward"`
	}

	r := req{MethodParams: params}

	var result models.PendingReward
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetPendingByNode(ctx context.Context, params GetPendingByNodeParams) (*models.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeParams `json:"get_pending_node_operator_reward"`
	}

	r := req{MethodParams: params}

	var result models.PendingReward
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetPendingByNodeAndDelegator(ctx context.Context, params GetPendingByNodeAndDelegatorParams) (*models.PendingReward, error) {
	type req struct {
		MethodParams GetPendingByNodeAndDelegatorParams `json:"get_pending_delegator_reward"`
	}

	r := req{MethodParams: params}

	var result models.PendingReward
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) EstimateOperatorReward(ctx context.Context, params EstimateOperatorRewardParams) (*models.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateOperatorRewardParams `json:"get_estimated_current_epoch_operator_reward"`
	}

	r := req{MethodParams: params}

	var result models.EstimatedCurrentEpochReward
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) EstimateDelegatorReward(ctx context.Context, params EstimateDelegatorRewardParams) (*models.EstimatedCurrentEpochReward, error) {
	type req struct {
		MethodParams EstimateDelegatorRewardParams `json:"get_estimated_current_epoch_delegator_reward"`
	}

	r := req{MethodParams: params}

	var result models.EstimatedCurrentEpochReward
	return &result, c.reader.Read(ctx, r, &result)
}
