package contract

import (
	"context"
	"github.com/craftdome/go-nym/pkg/contracts/mixnet/models"
)

func (c *Client) GetAdmin(ctx context.Context) (models.ContractAdmin, error) {
	type req struct {
		MethodParams struct{} `json:"admin"`
	}

	r := req{}

	return Query[models.ContractAdmin](ctx, c.client, c.contract, r)
}

func (c *Client) GetVersion(ctx context.Context) (models.ContractVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_contract_version"`
	}

	r := req{}

	return Query[models.ContractVersion](ctx, c.client, c.contract, r)
}

func (c *Client) GetCW2Version(ctx context.Context) (models.ContractCW2Version, error) {
	type req struct {
		MethodParams struct{} `json:"get_cw2_contract_version"`
	}

	r := req{}

	return Query[models.ContractCW2Version](ctx, c.client, c.contract, r)
}

func (c *Client) GetStateParams(ctx context.Context) (models.ContractStateParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_state_params"`
	}

	r := req{}

	return Query[models.ContractStateParams](ctx, c.client, c.contract, r)
}

func (c *Client) GetState(ctx context.Context) (models.ContractState, error) {
	type req struct {
		MethodParams struct{} `json:"get_state"`
	}

	r := req{}

	return Query[models.ContractState](ctx, c.client, c.contract, r)
}

func (c *Client) GetIntervalStatus(ctx context.Context) (models.IntervalStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_interval_details"`
	}

	r := req{}

	return Query[models.IntervalStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetEpochStatus(ctx context.Context) (models.EpochStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_epoch_status"`
	}

	r := req{}

	return Query[models.EpochStatus](ctx, c.client, c.contract, r)
}

func (c *Client) GetCurrentNodeVersion(ctx context.Context) (models.NodeVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_nym_node_version"`
	}

	r := req{}

	type CurrentNodeVersion struct {
		Version models.NodeVersion `json:"version"`
	}

	ver, err := Query[CurrentNodeVersion](ctx, c.client, c.contract, r)
	if err != nil {
		return models.NodeVersion{}, err
	}

	return ver.Version, nil
}

func (c *Client) GetNodeVersionHistory(ctx context.Context) (models.PagedNodeVersionHistory, error) {
	type req struct {
		MethodParams struct{} `json:"get_nym_node_version_history"`
	}

	r := req{}

	return Query[models.PagedNodeVersionHistory](ctx, c.client, c.contract, r)
}

func (c *Client) GetRewardingParams(ctx context.Context) (models.RewardingParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarding_params"`
	}

	r := req{}

	return Query[models.RewardingParams](ctx, c.client, c.contract, r)
}
