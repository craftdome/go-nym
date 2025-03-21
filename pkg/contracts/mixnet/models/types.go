package models

import (
	"fmt"
	"strconv"
	"uint128"
)

type NodeID uint32

func (n NodeID) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

type BlockHeight uint64
type IdentityKey string

type ProfitMarginRange struct {
	Minimum Percent `json:"minimum"`
	Maximum Percent `json:"maximum"`
}

func (r ProfitMarginRange) String() string {
	return fmt.Sprintf("[%s..=%s]", r.Minimum.String(), r.Maximum.String())
}

type OperatingCostRange struct {
	Minimum uint128.Uint128 `json:"minimum"`
	Maximum uint128.Uint128 `json:"maximum"`
}

func (r OperatingCostRange) String() string {
	return fmt.Sprintf("[%s..=%s]", r.Minimum.String(), r.Maximum.String())
}
