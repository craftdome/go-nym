package rewards

import "github.com/craftdome/go-nym/pkg/mixnet"

type GetPendingByOwnerParams struct {
	Owner mixnet.Addr `json:"address"`
}

type GetPendingByNodeParams struct {
	NodeID mixnet.NodeID `json:"node_id"`
}

type GetPendingByNodeAndDelegatorParams struct {
	Delegator mixnet.Addr   `json:"address"`
	NodeID    mixnet.NodeID `json:"node_id"`
}

type EstimateOperatorRewardParams struct {
	NodeID               mixnet.NodeID `json:"node_id"`
	EstimatedPerformance float32       `json:"estimated_performance,string"`
	EstimatedWork        float32       `json:"estimated_work,string"`
}

type EstimateDelegatorRewardParams struct {
	Delegator            mixnet.Addr   `json:"address"`
	NodeID               mixnet.NodeID `json:"node_id"`
	EstimatedPerformance float32       `json:"estimated_performance,string"`
	EstimatedWork        float32       `json:"estimated_work,string"`
}
