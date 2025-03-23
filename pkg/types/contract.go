package types

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type ContractAdmin struct {
	Admin string `json:"admin"`
}

type ContractVersion struct {
	ContractName    string `json:"contract_name"`
	BuildTimestamp  string `json:"build_timestamp"`
	BuildVersion    string `json:"build_version"`
	CommitSHA       string `json:"commit_sha"`
	CommitTimestamp string `json:"commit_timestamp"`
	CommitBranch    string `json:"commit_branch"`
	RustcVersion    string `json:"rustc_version"`
	CargoDebug      string `json:"cargo_debug"`
	CargoOptLevel   string `json:"cargo_opt_level"`
}

type ContractCW2Version struct {
	Contract string `json:"contract"`
	Version  string `json:"version"`
}

type ContractState struct {
	Owner                     Addr                `json:"owner"`
	RewardingValidatorAddress Addr                `json:"rewarding_validator_address"`
	VestingContractAddress    Addr                `json:"vesting_contract_address"`
	RewardingDenom            string              `json:"rewarding_denom"`
	Params                    ContractStateParams `json:"params"`
}

type ContractStateParams struct {
	DelegationParams struct {
		MinimumDelegation Coin `json:"minimum_delegation,omitzero"`
	} `json:"delegation_params"`
	OperatorsParams struct {
		MinimumPledge         Coin               `json:"minimum_pledge"`
		ProfitMargin          ProfitMarginRange  `json:"profit_margin"`
		IntervalOperatingCost OperatingCostRange `json:"interval_operating_cost"`
	} `json:"operators_params"`
	ConfigScoreParams struct {
		VersionWeights struct {
			Major      uint32 `json:"major"`
			Minor      uint32 `json:"minor"`
			Patch      uint32 `json:"patch"`
			Prerelease uint32 `json:"prerelease"`
		} `json:"version_weights"`
		VersionScoreFormulaParams struct {
			Penalty        float32 `json:"penalty,string"`
			PenaltyScaling float32 `json:"penalty_scaling,string"`
		} `json:"version_score_formula_params"`
	} `json:"config_score_params"`
}

type EpochID = uint32

type IntervalID = uint32

type OffsetDateTime struct{ time.Time }

func (dt *OffsetDateTime) UnmarshalText(text []byte) error {
	t, err := time.Parse(time.RFC3339, string(text))
	if err != nil {
		return err
	}

	*dt = OffsetDateTime{t}
	return nil
}

func (dt OffsetDateTime) String() string {
	return dt.Format(time.RFC3339)
}

type EpochStatus struct {
	BeingAdvancedBy Addr       `json:"being_advanced_by"`
	State           EpochState `json:"state"`
}

type InProgressEpochState = string

type RewardingEpochState struct {
	LastRewarded NodeID `json:"last_rewarded"`
	FinalNodeID  NodeID `json:"final_node_id"`
}
type ReconcilingEventsEpochState = string

type RoleAssignmentEpochState struct {
	Next Role `json:"next"`
}

type EpochState struct {
	*InProgressEpochState
	*RewardingEpochState
	*ReconcilingEventsEpochState
	*RoleAssignmentEpochState
}

func (es *EpochState) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return errors.New("empty epoch state")
	}

	// Проверяем, является ли входная data строкой
	if text[0] != '{' {
		s := string(text)

		// Обрабатываем строковое представление состояния
		switch s {
		case "in_progress":
			es.InProgressEpochState = &s
		case "reconciling_events":
			es.ReconcilingEventsEpochState = &s
		default:
			return errors.New("unknown EpochState: " + string(text))
		}
		return nil
	}

	// Если не строка, то парсим как объект
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(text, &raw); err != nil {
		return err
	}

	// Проверяем возможные варианты
	if v, ok := raw["rewarding"]; ok {
		return json.Unmarshal(v, &es.RewardingEpochState)
	}
	if v, ok := raw["role_assignment"]; ok {
		return json.Unmarshal(v, &es.RoleAssignmentEpochState)
	}

	return errors.New("unknown EpochState")
}

func (es EpochState) String() string {
	switch {
	case es.InProgressEpochState != nil:
		return fmt.Sprintf("%+v", es.InProgressEpochState)
	case es.RewardingEpochState != nil:
		return fmt.Sprintf("%+v", es.RewardingEpochState)
	case es.ReconcilingEventsEpochState != nil:
		return fmt.Sprintf("%+v", es.ReconcilingEventsEpochState)
	case es.RoleAssignmentEpochState != nil:
		return fmt.Sprintf("%+v", es.RoleAssignmentEpochState)
	default:
		return "<nil>"
	}
}

type Interval struct {
	ID                IntervalID     `json:"id"`
	EpochsInInterval  uint32         `json:"epochs_in_interval"`
	CurrentEpochStart OffsetDateTime `json:"current_epoch_start"`
	EpochLength       struct {
		Secs  uint32 `json:"secs"`
		Nanos uint32 `json:"nanos"`
	} `json:"epoch_length"`
	CurrentEpochID     EpochID `json:"current_epoch_id"`
	TotalElapsedEpochs EpochID `json:"total_elapsed_epochs"`
}

type IntervalStatus struct {
	Interval              Interval `json:"interval"`
	CurrentBlocktime      uint64   `json:"current_blocktime"`
	IsCurrentIntervalOver bool     `json:"is_current_interval_over"`
	IsCurrentEpochOver    bool     `json:"is_current_epoch_over"`
}

type PagedNodeVersionHistory struct {
	History        []NodeVersion `json:"history"`
	StartNextAfter uint32        `json:"start_next_after,omitempty"`
}

type NodeVersion struct {
	ID   uint32          `json:"id"`
	Info NodeVersionInfo `json:"version_information"`
}

type NodeVersionInfo struct {
	Semver                 string `json:"semver"`
	IntroducedAtHeight     uint64 `json:"introduced_at_height"`
	DifferenceSinceGenesis struct {
		Major      uint32 `json:"major"`
		Minor      uint32 `json:"minor"`
		Patch      uint32 `json:"patch"`
		Prerelease uint32 `json:"prerelease"`
	} `json:"difference_since_genesis"`
}
