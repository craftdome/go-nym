package mixnet

import (
	"context"
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm/types"
)

type client struct {
	client   types.QueryClient
	contract string
}

func (c *client) Read(ctx context.Context, query, dest any) error {
	data, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	req := types.QuerySmartContractStateRequest{
		Address:   c.contract,
		QueryData: data,
	}

	resp, err := c.client.SmartContractState(ctx, &req)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Data, dest); err != nil {
		return err
	}

	return nil
}
