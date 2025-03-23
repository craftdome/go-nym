package contract

import (
	"context"
	"github.com/craftdome/go-nym/pkg/types"
)

func (c *Client) GetAdmin(ctx context.Context) (types.ContractAdmin, error) {
	type req struct {
		MethodParams struct{} `json:"admin"`
	}

	r := req{}

	return Query[types.ContractAdmin](ctx, c.client, c.contract, r)
}

func (c *Client) GetVersion(ctx context.Context) (types.ContractVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_contract_version"`
	}

	r := req{}

	return Query[types.ContractVersion](ctx, c.client, c.contract, r)
}

func (c *Client) GetCW2Version(ctx context.Context) (types.ContractCW2Version, error) {
	type req struct {
		MethodParams struct{} `json:"get_cw2_contract_version"`
	}

	r := req{}

	return Query[types.ContractCW2Version](ctx, c.client, c.contract, r)
}

func (c *Client) GetStateParams(ctx context.Context) (types.ContractStateParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_state_params"`
	}

	r := req{}

	return Query[types.ContractStateParams](ctx, c.client, c.contract, r)
}

func (c *Client) GetState(ctx context.Context) (types.ContractState, error) {
	type req struct {
		MethodParams struct{} `json:"get_state"`
	}

	r := req{}

	return Query[types.ContractState](ctx, c.client, c.contract, r)
}

func (c *Client) GetIntervalStatus(ctx context.Context) (types.IntervalStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_interval_details"`
	}

	r := req{}

	return Query[types.IntervalStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochStatus(ctx context.Context) (types.EpochStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_epoch_status"`
	}

	r := req{}

	return Query[types.EpochStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetCurrentNodeVersion(ctx context.Context) (types.NodeVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_nym_node_version"`
	}

	r := req{}

	type CurrentNodeVersion struct {
		Version types.NodeVersion `json:"version"`
	}

	ver, err := Query[CurrentNodeVersion](ctx, c.client, c.contract, r)
	if err != nil {
		return types.NodeVersion{}, err
	}

	return ver.Version, nil
}

func (c *Client) GetNodeVersionHistory(ctx context.Context) (types.PagedNodeVersionHistory, error) {
	type req struct {
		MethodParams struct{} `json:"get_nym_node_version_history"`
	}

	r := req{}

	return Query[types.PagedNodeVersionHistory](ctx, c.client, c.contract, r)
}

func (c *Client) GetRewardingParams(ctx context.Context) (types.RewardingParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarding_params"`
	}

	r := req{}

	return Query[types.RewardingParams](ctx, c.client, c.contract, r)
}
