package mixnet

type Performance = Percent

type WorkFactor = Decimal

type RewardingParams struct {
	Interval struct {
		RewardPool               Decimal `json:"reward_pool"`
		StakingSupply            Decimal `json:"staking_supply"`
		StakingSupplyScaleFactor Percent `json:"staking_supply_scale_factor"`
		EpochRewardBudget        Decimal `json:"epoch_reward_budget"`
		StakeSaturationPoint     Decimal `json:"stake_saturation_point"`
		SybilResistance          Percent `json:"sybil_resistance"`
		ActiveSetWorkFactor      Decimal `json:"active_set_work_factor"`
		IntervalPoolEmission     Percent `json:"interval_pool_emission"`
	} `json:"interval"`
	RewardedSet RewardedSetParams `json:"rewarded_set"`
}

type RewardedSetParams struct {
	EntryGateways uint32 `json:"entry_gateways"`
	ExitGateways  uint32 `json:"exit_gateways"`
	Mixnodes      uint32 `json:"mixnodes"`
	Standby       uint32 `json:"standby"`
}

type NodeRewardingParameters struct {
	// Performance of the particular node in the current epoch.
	Performance Performance `json:"performance"`

	// Amount of work performed by this node in the current epoch
	// also known as 'omega' in the paper
	WorkFactor WorkFactor `json:"work_factor"`
}

type IntervalRewardingParamsUpdate struct {
	RewardPool               Decimal           `json:"reward_pool"`
	StakingSupply            Decimal           `json:"staking_supply"`
	StakingSupplyScaleFactor Percent           `json:"staking_supply_scale_factor"`
	SybilResistancePercent   Percent           `json:"sybil_resistance_percent"`
	ActiveSetWorkFactor      Decimal           `json:"active_set_work_factor"`
	IntervalPoolEmission     Percent           `json:"interval_pool_emission"`
	RewardedSetParams        RewardedSetParams `json:"rewarded_set_params"`
}

type ActiveSetUpdate struct {
	EntryGateways uint32 `json:"entry_gateways"`
	ExitGateways  uint32 `json:"exit_gateways"`
	Mixnodes      uint32 `json:"mixnodes"`
}

// PendingReward https://github.com/nymtech/nym/blob/develop/common/cosmwasm-smart-contracts/mixnet-contract/src/rewarding/mod.rs#L60
type PendingReward struct {
	AmountStaked         Coin    `json:"amount_staked"`
	AmountEarned         Coin    `json:"amount_earned"`
	AmountEarnedDetailed Decimal `json:"amount_earned_detailed"`
	NodeStillFullyBonded bool    `json:"node_still_fully_bonded"`
}

// RewardDistribution https://github.com/nymtech/nym/blob/develop/common/cosmwasm-smart-contracts/mixnet-contract/src/rewarding/mod.rs#L60
type RewardDistribution struct {
	Operator  Decimal `json:"operator"`
	Delegates Decimal `json:"delegates"`
}

// EstimatedCurrentEpochReward https://github.com/nymtech/nym/blob/9c4243914e68b58b7288c75494e9a6774e8b3601/common/cosmwasm-smart-contracts/mixnet-contract/src/rewarding/mod.rs#L81
type EstimatedCurrentEpochReward struct {
	OriginalStake                   Coin    `json:"original_stake"`
	CurrentStakeValue               Coin    `json:"current_stake_value"`
	CurrentStakeValueDetailedAmount Decimal `json:"current_stake_value_detailed_amount"`
	Estimation                      Coin    `json:"estimation"`
	DetailedEstimationAmount        Decimal `json:"detailed_estimation_amount"`
}
