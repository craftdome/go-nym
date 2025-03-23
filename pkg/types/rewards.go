package types

type RewardingParams struct {
	Interval struct {
		RewardPool               Uint64  `json:"reward_pool"`
		StakingSupply            Uint64  `json:"staking_supply"`
		StakingSupplyScaleFactor float32 `json:"staking_supply_scale_factor,string"`
		EpochRewardBudget        Uint64  `json:"epoch_reward_budget"`
		StakeSaturationPoint     Uint64  `json:"stake_saturation_point"`
		SybilResistance          float32 `json:"sybil_resistance,string"`
		ActiveSetWorkFactor      float32 `json:"active_set_work_factor,string"`
		IntervalPoolEmission     float32 `json:"interval_pool_emission,string"`
	} `json:"interval"`
	RewardedSet RewardedSetParams `json:"rewarded_set"`
}

type RewardedSetParams struct {
	EntryGateways uint32 `json:"entry_gateways"`
	ExitGateways  uint32 `json:"exit_gateways"`
	Mixnodes      uint32 `json:"mixnodes"`
	Standby       uint32 `json:"standby"`
}

type IntervalRewardingParamsUpdate struct {
	RewardPool               Uint64            `json:"reward_pool"`
	StakingSupply            Uint64            `json:"staking_supply"`
	StakingSupplyScaleFactor float32           `json:"staking_supply_scale_factor,string"`
	SybilResistancePercent   float32           `json:"sybil_resistance_percent,string"`
	ActiveSetWorkFactor      float32           `json:"active_set_work_factor,string"`
	IntervalPoolEmission     float32           `json:"interval_pool_emission,string"`
	RewardedSetParams        RewardedSetParams `json:"rewarded_set_params"`
}

type ActiveSetUpdate struct {
	EntryGateways uint32 `json:"entry_gateways"`
	ExitGateways  uint32 `json:"exit_gateways"`
	Mixnodes      uint32 `json:"mixnodes"`
}

// PendingReward https://github.com/nymtech/nym/blob/develop/common/cosmwasm-smart-contracts/mixnet-contract/src/rewarding/mod.rs#L60
type PendingReward struct {
	AmountStaked         Coin   `json:"amount_staked"`
	AmountEarned         Coin   `json:"amount_earned"`
	AmountEarnedDetailed Uint64 `json:"amount_earned_detailed"`
	NodeStillFullyBonded bool   `json:"node_still_fully_bonded"`
}

// EstimatedCurrentEpochReward https://github.com/nymtech/nym/blob/9c4243914e68b58b7288c75494e9a6774e8b3601/common/cosmwasm-smart-contracts/mixnet-contract/src/rewarding/mod.rs#L81
type EstimatedCurrentEpochReward struct {
	OriginalStake     Coin `json:"original_stake"`
	CurrentStakeValue Coin `json:"current_stake_value"`
	Estimation        Coin `json:"estimation"`
}
