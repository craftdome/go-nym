package contract

import (
	"context"
	"github.com/craftdome/go-nym/pkg/mixnet"
)

func (c *Client) GetAdmin(ctx context.Context) (mixnet.ContractAdmin, error) {
	type req struct {
		MethodParams struct{} `json:"admin"`
	}

	r := req{}

	return Query[mixnet.ContractAdmin](ctx, c.client, c.contract, r)
}

func (c *Client) GetVersion(ctx context.Context) (mixnet.ContractVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_contract_version"`
	}

	r := req{}

	return Query[mixnet.ContractVersion](ctx, c.client, c.contract, r)
}

func (c *Client) GetCW2Version(ctx context.Context) (mixnet.ContractCW2Version, error) {
	type req struct {
		MethodParams struct{} `json:"get_cw2_contract_version"`
	}

	r := req{}

	return Query[mixnet.ContractCW2Version](ctx, c.client, c.contract, r)
}

func (c *Client) GetStateParams(ctx context.Context) (mixnet.ContractStateParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_state_params"`
	}

	r := req{}

	return Query[mixnet.ContractStateParams](ctx, c.client, c.contract, r)
}

func (c *Client) GetState(ctx context.Context) (mixnet.ContractState, error) {
	type req struct {
		MethodParams struct{} `json:"get_state"`
	}

	r := req{}

	return Query[mixnet.ContractState](ctx, c.client, c.contract, r)
}

func (c *Client) GetIntervalStatus(ctx context.Context) (mixnet.IntervalStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_interval_details"`
	}

	r := req{}

	return Query[mixnet.IntervalStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochStatus(ctx context.Context) (mixnet.EpochStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_epoch_status"`
	}

	r := req{}

	return Query[mixnet.EpochStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetCurrentNodeVersion(ctx context.Context) (mixnet.NodeVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_nym_node_version"`
	}

	r := req{}

	type CurrentNodeVersion struct {
		Version mixnet.NodeVersion `json:"version"`
	}

	ver, err := Query[CurrentNodeVersion](ctx, c.client, c.contract, r)
	if err != nil {
		return mixnet.NodeVersion{}, err
	}

	return ver.Version, nil
}

func (c *Client) GetNodeVersionHistory(ctx context.Context) (mixnet.PagedNodeVersionHistory, error) {
	type req struct {
		MethodParams struct{} `json:"get_nym_node_version_history"`
	}

	r := req{}

	return Query[mixnet.PagedNodeVersionHistory](ctx, c.client, c.contract, r)
}

func (c *Client) GetRewardingParams(ctx context.Context) (mixnet.RewardingParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarding_params"`
	}

	r := req{}

	return Query[mixnet.RewardingParams](ctx, c.client, c.contract, r)
}
