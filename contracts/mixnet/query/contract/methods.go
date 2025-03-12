package contract

import (
	"context"
	"github.com/craftdome/go-nym/contracts/mixnet/shared/models"
)

func (c *Client) GetAdmin(ctx context.Context) (*models.ContractAdmin, error) {
	type req struct {
		MethodParams struct{} `json:"admin"`
	}

	r := req{}

	var result models.ContractAdmin
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetVersion(ctx context.Context) (*models.ContractVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_contract_version"`
	}

	r := req{}

	var result models.ContractVersion
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetCW2Version(ctx context.Context) (*models.ContractCW2Version, error) {
	type req struct {
		MethodParams struct{} `json:"get_cw2_contract_version"`
	}

	r := req{}

	var result models.ContractCW2Version
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetStateParams(ctx context.Context) (*models.ContractStateParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_state_params"`
	}

	r := req{}

	var result models.ContractStateParams
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetState(ctx context.Context) (*models.ContractState, error) {
	type req struct {
		MethodParams struct{} `json:"get_state"`
	}

	r := req{}

	var result models.ContractState
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetIntervalStatus(ctx context.Context) (*models.IntervalStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_interval_details"`
	}

	r := req{}

	var result models.IntervalStatus
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetEpochStatus(ctx context.Context) (*models.EpochStatus, error) {
	type req struct {
		MethodParams struct{} `json:"get_epoch_status"`
	}

	r := req{}

	var result models.EpochStatus
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetCurrentNodeVersion(ctx context.Context) (*models.NodeVersion, error) {
	type req struct {
		MethodParams struct{} `json:"get_current_nym_node_version"`
	}

	r := req{}

	type resp struct {
		Version models.NodeVersion `json:"version"`
	}

	var result resp
	return &result.Version, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetNodeVersionHistory(ctx context.Context) (*models.PagedNodeVersionHistory, error) {
	type req struct {
		MethodParams struct{} `json:"get_nym_node_version_history"`
	}

	r := req{}

	var result models.PagedNodeVersionHistory
	return &result, c.reader.Read(ctx, r, &result)
}

func (c *Client) GetRewardingParams(ctx context.Context) (*models.RewardingParams, error) {
	type req struct {
		MethodParams struct{} `json:"get_rewarding_params"`
	}

	r := req{}

	var result models.RewardingParams
	return &result, c.reader.Read(ctx, r, &result)
}
