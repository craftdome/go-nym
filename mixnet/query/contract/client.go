package contract

import (
	"context"
	"encoding/json"
	"github.com/CosmWasm/wasmd/x/wasm/types"
)

type Client struct {
	client   types.QueryClient
	contract string
}

func New(client types.QueryClient, contract string) *Client {
	return &Client{
		client:   client,
		contract: contract,
	}
}

func Query[T any](ctx context.Context, client types.QueryClient, contract string, query any) (T, error) {
	var result T
	data, err := json.Marshal(&query)
	if err != nil {
		return result, err
	}

	req := types.QuerySmartContractStateRequest{
		Address:   contract,
		QueryData: data,
	}

	resp, err := client.SmartContractState(ctx, &req)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return result, err
	}

	return result, nil
}

func QueryRaw[T any](ctx context.Context, client types.QueryClient, contract string, queryData []byte) (T, error) {
	var result T

	req := types.QuerySmartContractStateRequest{
		Address:   contract,
		QueryData: queryData,
	}

	resp, err := client.SmartContractState(ctx, &req)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return result, err
	}

	return result, nil
}
