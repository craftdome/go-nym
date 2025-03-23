package types_test

import (
	"github.com/craftdome/go-nym/pkg/types"
	"math"
	"testing"
)

func TestUint64_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    types.Uint64
		wantErr bool
	}{
		{
			name:    "integer value",
			input:   "123456789",
			want:    types.Uint64(123456789),
			wantErr: false,
		},
		{
			name:    "with fractional part",
			input:   "332441515.4838589327582",
			want:    types.Uint64(332441515),
			wantErr: false,
		},
		{
			name:    "rounding down to zero",
			input:   "0.999",
			want:    types.Uint64(0),
			wantErr: false,
		},
		{
			name:    "maximum uint64 value",
			input:   "18446744073709551615", // max uint64
			want:    types.Uint64(math.MaxUint64),
			wantErr: false,
		},
		{
			name:    "overflow uint64",
			input:   "18446744073709551616", // max uint64 + 1
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid input",
			input:   "abc.def",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u types.Uint64
			err := u.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if u != tt.want {
				t.Errorf("UnmarshalText() = %v, want = %v", u, tt.want)
			}
		})
	}
}
