package types

import (
	"github.com/pkg/errors"
	"math/big"
)

type Uint64 uint64

func (u *Uint64) UnmarshalText(text []byte) error {
	f, _, err := big.ParseFloat(string(text), 10, 0, big.ToZero)
	if err != nil {
		return err
	}

	i, _ := f.Int(&big.Int{})
	if !i.IsUint64() {
		return errors.New("value out of range")
	}

	*u = Uint64(i.Uint64())
	return nil
}
