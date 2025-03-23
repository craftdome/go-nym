package types

type NodeRewardingDetails struct {
	CostParams        NodeCostParams `json:"cost_params"`
	Operator          Uint64         `json:"operator"`
	Delegates         Uint64         `json:"delegates"`
	TotalUnitReward   Uint64         `json:"total_unit_reward"`
	UnitDelegation    Uint64         `json:"unit_delegation"`
	LastRewardedEpoch EpochID        `json:"last_rewarded_epoch"`
	UniqueDelegations uint32         `json:"unique_delegations"`
}

type NodeCostParams struct {
	ProfitMarginPercent   Percent `json:"profit_margin_percent,string"`
	IntervalOperatingCost Coin    `json:"interval_operating_cost"`
}

type Role = string

const (
	EntryGateway Role = "eg"
	Layer1       Role = "l1"
	Layer2       Role = "l2"
	Layer3       Role = "l3"
	ExitGateway  Role = "xg"
	Standby      Role = "stb"
)

type RewardedSetMetadata struct {
	EpochID              EpochID      `json:"epoch_id"`
	FullyAssigned        bool         `json:"fully_assigned"`
	EntryGatewayMetadata RoleMetadata `json:"entry_gateway_metadata"`
	ExitGatewayMetadata  RoleMetadata `json:"exit_gateway_metadata"`
	Layer1Metadata       RoleMetadata `json:"layer1_metadata"`
	Layer2Metadata       RoleMetadata `json:"layer2_metadata"`
	Layer3Metadata       RoleMetadata `json:"layer3_metadata"`
	StandbyMetadata      RoleMetadata `json:"standby_metadata"`
}

type RoleMetadata struct {
	HighestID NodeID `json:"highest_id"`
	NumNodes  uint32 `json:"num_nodes"`
}

type DetailedNode struct {
	BondInformation  BondedNode           `json:"bond_information"`
	RewardingDetails NodeRewardingDetails `json:"rewarding_details"`
	PendingChanges   PendingNodeChanges   `json:"pending_changes"`
}

type BondedNode struct {
	NodeID         NodeID `json:"node_id"`
	Owner          Addr   `json:"owner"`
	OriginalPledge Coin   `json:"original_pledge"`
	BondingHeight  uint64 `json:"bonding_height"`
	IsUnbonding    bool   `json:"is_unbonding"`
	Node           Node   `json:"node"`
}

type Node struct {
	Host           string      `json:"host"`
	CustomHTTPPort uint16      `json:"custom_http_port"`
	IdentityKey    IdentityKey `json:"identity_key"`
}

type NodeConfigUpdate struct {
	Host                   string `json:"host,omitempty"`
	CustomHTTPPort         uint16 `json:"custom_http_port"`
	RestoreDefaultHTTPPort bool   `json:"restore_default_http_port"`
}

type PendingNodeChanges struct {
	PledgeChange     *EpochEventID    `json:"pledge_change"`
	CostParamsChange *IntervalEventID `json:"cost_params_change"`
}

type UnbondedNode struct {
	IdentityKey     IdentityKey `json:"identity_key"`
	NodeID          NodeID      `json:"node_id"`
	Operator        Addr        `json:"owner"`
	UnbondingHeight uint64      `json:"unbonding_height"`
}

type DetailedByOperator struct {
	Operator Addr          `json:"address"`
	Details  *DetailedNode `json:"details"`
}

type NodeStakeSaturation struct {
	NodeID             NodeID  `json:"node_id"`
	CurrentSaturation  Percent `json:"current_saturation,string"`
	UncappedSaturation Percent `json:"uncapped_saturation,string"`
}

type PagedUnbondedNodes struct {
	Nodes          []UnbondedNode `json:"nodes"`
	StartNextAfter NodeID         `json:"start_next_after"`
}

type PagedBondedNodes struct {
	Nodes          []BondedNode `json:"nodes"`
	StartNextAfter NodeID       `json:"start_next_after"`
}

type PagedDetailedNodes struct {
	Nodes          []DetailedNode `json:"nodes"`
	StartNextAfter NodeID         `json:"start_next_after"`
}

type EpochAssignment struct {
	EpochID EpochID  `json:"epoch_id"`
	Nodes   []NodeID `json:"nodes"`
}

type EpochAssignmentMetadata struct {
	Metadata RewardedSetMetadata `json:"metadata"`
}
