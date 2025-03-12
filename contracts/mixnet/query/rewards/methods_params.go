package rewards

import (
	"contracts/mixnet/shared/models"
)

type GetPendingByOwnerParams struct {
	Owner models.Addr `json:"address"`
}

type GetPendingByNodeParams struct {
	NodeID models.NodeID `json:"node_id"`
}

type GetPendingByNodeAndDelegatorParams struct {
	Delegator models.Addr   `json:"address"`
	NodeID    models.NodeID `json:"node_id"`
}

type EstimateOperatorRewardParams struct {
	NodeID               models.NodeID `json:"node_id"`
	EstimatedPerformance float32       `json:"estimated_performance,string"`
	EstimatedWork        float32       `json:"estimated_work,string"`
}

type EstimateDelegatorRewardParams struct {
	Delegator            models.Addr   `json:"address"`
	NodeID               models.NodeID `json:"node_id"`
	EstimatedPerformance float32       `json:"estimated_performance,string"`
	EstimatedWork        float32       `json:"estimated_work,string"`
}
