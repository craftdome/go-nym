package types

import (
	"fmt"
)

type NodeID = uint32
type BlockHeight = uint64
type IdentityKey = string

type ProfitMarginRange struct {
	Minimum Percent `json:"minimum,string"`
	Maximum Percent `json:"maximum,string"`
}

func (r ProfitMarginRange) String() string {
	return fmt.Sprintf("[%.2f..=%.2f]", r.Minimum, r.Maximum)
}

type OperatingCostRange struct {
	Minimum uint64 `json:"minimum,string"`
	Maximum uint64 `json:"maximum,string"`
}

func (r OperatingCostRange) String() string {
	return fmt.Sprintf("[%d..=%d]", r.Minimum, r.Maximum)
}
