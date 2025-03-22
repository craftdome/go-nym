package mixnet

import (
	"github.com/craftdome/go-nym/pkg/uint128"
	"math/big"
)

var (
	TokenSupply                        = uint128.From64(1_000_000_000_000_000)
	DefaultIntervalOperatingCostAmount = uint128.From64(40_000_000)
	DefaultProfitMarginPercent         = uint64(20)
	UnitDelegationBase                 = new(big.Int).Mul(big.NewInt(1_000_000_000), big.NewInt(1_000_000_000_000_000_000))
)

const (
	NymNodeBondDefaultRetrievalLimit uint32 = 50
	NymNodeBondMaxRetrievalLimit     uint32 = 100

	NymNodeDetailsDefaultRetrievalLimit uint32 = 50
	NymNodeDetailsMaxRetrievalLimit     uint32 = 75

	UnbondedNymNodesDefaultRetrievalLimit uint32 = 100
	UnbondedNymNodesMaxRetrievalLimit     uint32 = 200

	DelegationPageDefaultRetrievalLimit uint32 = 100
	DelegationPageMaxRetrievalLimit     uint32 = 500

	EpochEventsDefaultRetrievalLimit uint32 = 50
	EpochEventsMaxRetrievalLimit     uint32 = 100

	IntervalEventsDefaultRetrievalLimit uint32 = 50
	IntervalEventsMaxRetrievalLimit     uint32 = 100
)
