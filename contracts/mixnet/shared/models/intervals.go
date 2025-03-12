package models

type IntervalEventID uint32

func (id IntervalEventID) IsZero() bool {
	return id == 0
}

type EpochEventID uint32

func (id EpochEventID) IsZero() bool {
	return id == 0
}

type PendingEpochEvent struct {
	EventID EpochEventID          `json:"event_id"`
	Event   PendingEpochEventData `json:"event"`
}

type PendingEpochEventData struct {
	CreatedAt BlockHeight           `json:"created_at"`
	Kind      PendingEpochEventKind `json:"kind"`
}

type PendingEpochEventKind struct {
	Delegate *struct {
		Owner  Addr   `json:"owner"`
		NodeID NodeID `json:"node_id"`
		Amount Coin   `json:"amount"`
		Proxy  *Addr  `json:"proxy"`
	} `json:"delegate"`
	Undelegate *struct {
		Owner  Addr   `json:"owner"`
		NodeID NodeID `json:"node_id"`
		Proxy  *Addr  `json:"proxy"`
	} `json:"undelegate"`
	NodeIncreasePledge *struct {
		NodeID     NodeID `json:"node_id"`
		IncreaseBy Coin   `json:"amount"`
	} `json:"nym_node_pledge_more"`
	MixnodePledgeMore *struct {
		MixID  NodeID `json:"mix_id"`
		Amount Coin   `json:"amount"`
	} `json:"mixnode_pledge_more"`
	NodeDecreasePledge *struct {
		NodeID     NodeID `json:"node_id"`
		DecreaseBy Coin   `json:"decrease_by"`
	} `json:"nym_node_decrease_pledge"`
	MixnodeDecreasePledge *struct {
		MixID      NodeID `json:"mix_id"`
		DecreaseBy Coin   `json:"decrease_by"`
	} `json:"mixnode_decrease_pledge"`
	UnbondMixnode *struct {
		MixID NodeID `json:"mix_id"`
	} `json:"unbond_mixnode"`
	UnbondNode *struct {
		NodeID NodeID `json:"node_id"`
	} `json:"unbond_nym_node"`
	UpdateActiveSet *struct {
		Update ActiveSetUpdate `json:"update"`
	} `json:"update_active_set"`
}

type PendingIntervalEvent struct {
	EventID IntervalEventID          `json:"event_id"`
	Event   PendingIntervalEventData `json:"event"`
}

type PendingIntervalEventData struct {
	CreatedAt BlockHeight              `json:"created_at"`
	Kind      PendingIntervalEventKind `json:"kind"`
}

type PendingIntervalEventKind struct {
	ChangeMixCostParams *struct {
		MixID    NodeID         `json:"mix_id"`
		NewCosts NodeCostParams `json:"new_costs"`
	} `json:"change_mix_cost_params"`
	ChangeNodeCostParams *struct {
		NodeID   NodeID         `json:"node_id"`
		NewCosts NodeCostParams `json:"new_costs"`
	} `json:"change_nym_node_cost_params"`
	UpdateRewardingParams *struct {
		Update IntervalRewardingParamsUpdate `json:"update"`
	} `json:"update_rewarding_params"`
	UpdateIntervalConfig *struct {
		EpochsInInterval  uint32 `json:"epochs_in_interval"`
		EpochDurationSecs uint64 `json:"epoch_duration_secs"`
	} `json:"update_interval_config"`
}

type PendingEpochEvents struct {
	SecondsUntilExecutable int64 `json:"seconds_until_executable"`
	Events                 []struct {
		ID    EpochEventID          `json:"id"`
		Event PendingEpochEventData `json:"event"`
	} `json:"events"`
	StartNextAfter EpochEventID `json:"start_next_after"`
}

type PendingIntervalEvents struct {
	SecondsUntilExecutable int64 `json:"seconds_until_executable"`
	Events                 []struct {
		ID    IntervalEventID          `json:"id"`
		Event PendingIntervalEventData `json:"event"`
	} `json:"events"`
	StartNextAfter IntervalEventID `json:"start_next_after"`
}

type NumberOfPendingEvents struct {
	EpochEvents    uint32 `json:"epoch_events"`
	IntervalEvents uint32 `json:"interval_events"`
}
