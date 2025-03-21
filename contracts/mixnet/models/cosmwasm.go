package models

import (
	"math/big"
	"strconv"
	"uint128"
)

type Addr string

func (a Addr) IsZero() bool {
	return a == ""
}

type Coin struct {
	Denom  string          `json:"denom"`
	Amount uint128.Uint128 `json:"amount"`
}

func (c Coin) IsZero() bool {
	return c.Amount.IsZero()
}

type Decimal struct{ big.Float }

func (d Decimal) String() string {
	v, _ := d.Uint64()
	return strconv.FormatUint(v, 10)
}

func (d Decimal) IsZero() bool {
	v, _ := d.Float64()
	return v == 0
}

type Percent struct{ big.Float }

func (p Percent) String() string {
	v, _ := p.Float64()
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (p Percent) IsZero() bool {
	v, _ := p.Float64()
	return v == 0
}
