package rewards

import "github.com/craftdome/go-nym/pkg/types"

type GetPendingByOwnerParams struct {
	Owner types.Addr `json:"address"`
}

type GetPendingByNodeParams struct {
	NodeID types.NodeID `json:"node_id"`
}

type GetPendingByNodeAndDelegatorParams struct {
	Delegator types.Addr   `json:"address"`
	NodeID    types.NodeID `json:"node_id"`
}

type EstimateOperatorRewardParams struct {
	NodeID               types.NodeID `json:"node_id"`
	EstimatedPerformance float32      `json:"estimated_performance,string"`
	EstimatedWork        float32      `json:"estimated_work,string"`
}

type EstimateDelegatorRewardParams struct {
	Delegator            types.Addr   `json:"address"`
	NodeID               types.NodeID `json:"node_id"`
	EstimatedPerformance float32      `json:"estimated_performance,string"`
	EstimatedWork        float32      `json:"estimated_work,string"`
}
