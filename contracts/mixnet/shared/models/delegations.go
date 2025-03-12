package models

import "github.com/barweiss/go-tuple"

type OwnerProxySubKey Addr

type StorageKey = tuple.T2[NodeID, Addr]

type Delegation struct {
	Owner                 Addr    `json:"owner"`
	NodeID                NodeID  `json:"node_id"`
	CumulativeRewardRatio Decimal `json:"cumulative_reward_ratio"`
	Amount                Coin
	Height                uint64 `json:"height"`
	Proxy                 *Addr  `json:"proxy"`
}

type PagedNodeDelegations struct {
	Delegations    []Delegation      `json:"delegations"`
	StartNextAfter *OwnerProxySubKey `json:"start_next_after"`
}

type PagedDelegatorDelegations struct {
	Delegations    []Delegation `json:"delegations"`
	StartNextAfter *StorageKey  `json:"start_next_after"`
}

type DelegatorNodeDelegation struct {
	Delegation      *Delegation `json:"delegation"`
	NodeStillBonded bool        `json:"node_still_bonded"`
}

type PagedAllDelegations struct {
	Delegations    []Delegation `json:"delegations"`
	StartNextAfter *StorageKey  `json:"start_next_after"`
}
