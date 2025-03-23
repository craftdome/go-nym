package types

type Addr = string

type Coin struct {
	Denom  string `json:"denom"`
	Amount uint64 `json:"amount,string"`
}

type Percent = float32
