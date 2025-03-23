package types_test

import (
	"github.com/craftdome/go-nym/pkg/types"
	"reflect"
	"testing"
	"time"
)

func TestEpochState_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    types.EpochState
		wantErr bool
	}{
		{
			name:  "in_progress string",
			input: `in_progress`,
			want: types.EpochState{
				InProgressEpochState: ptr("in_progress"),
			},
			wantErr: false,
		},
		{
			name:  "reconciling_events string",
			input: `reconciling_events`,
			want: types.EpochState{
				ReconcilingEventsEpochState: ptr("reconciling_events"),
			},
			wantErr: false,
		},
		{
			name:  "rewarding object",
			input: `{"rewarding":{"last_rewarded":1,"final_node_id":2}}`,
			want: types.EpochState{
				RewardingEpochState: &types.RewardingEpochState{
					LastRewarded: 1,
					FinalNodeID:  2,
				},
			},
			wantErr: false,
		},
		{
			name:  "role_assignment object",
			input: `{"role_assignment":{"next":"l1"}}`,
			want: types.EpochState{
				RoleAssignmentEpochState: &types.RoleAssignmentEpochState{
					Next: types.Layer1,
				},
			},
			wantErr: false,
		},
		{
			name:    "unknown string",
			input:   `something_else`,
			want:    types.EpochState{},
			wantErr: true,
		},
		{
			name:    "unknown object field",
			input:   `{"unexpected_field":{"key":"value"}}`,
			want:    types.EpochState{},
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `{invalid}`,
			want:    types.EpochState{},
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   ``,
			want:    types.EpochState{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var es types.EpochState
			err := es.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(es, tt.want) {
				t.Errorf("UnmarshalText() got = %#v, want = %#v", es, tt.want)
			}
		})
	}
}

func ptr(s string) *string {
	return &s
}

func TestOffsetDateTime_UnmarshalText(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect time.Time
		err    bool
	}{
		{
			name:   "Valid time in UTC",
			input:  "2025-03-05T09:26:19Z",
			expect: time.Date(2025, time.March, 5, 9, 26, 19, 0, time.UTC),
			err:    false,
		},
		{
			name:   "Valid time with offset",
			input:  "2023-03-25T14:30:00+02:00",
			expect: time.Date(2023, time.March, 25, 12, 30, 0, 0, time.UTC),
			err:    false,
		},
		{
			name:  "Invalid time",
			input: "invalid-date",
			err:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var dt types.OffsetDateTime
			err := dt.UnmarshalText([]byte(tc.input))
			if tc.err {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Сравнение времени выполняется через Equal, которое учитывает абсолютное значение времени.
			if !dt.Time.Equal(tc.expect) {
				t.Errorf("Expected time %v, got %v", tc.expect, dt.Time)
			}
		})
	}
}
