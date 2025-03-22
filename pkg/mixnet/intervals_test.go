package mixnet

import (
	"reflect"
	"testing"
	"time"
)

func TestEpochState_UnmarshalText(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect EpochState
		err    bool
	}{
		{
			name:  "String in_progress",
			input: `in_progress`,
			expect: EpochState{
				InProgressEpochState: "in_progress",
				flag:                 inProgress,
			},
			err: false,
		},
		{
			name:  "String reconciling_events",
			input: `reconciling_events`,
			expect: EpochState{
				ReconcilingEventsEpochState: "reconciling_events",
				flag:                        reconcilingEvents,
			},
			err: false,
		},
		{
			name:  "Object rewarding",
			input: `{"rewarding": {"last_rewarded": 1, "final_node_id": 2}}`,
			expect: EpochState{
				RewardingEpochState: RewardingEpochState{
					LastRewarded: 1,
					FinalNodeID:  2,
				},
				flag: rewarding,
			},
			err: false,
		},
		{
			name:  "Object role_assignment",
			input: `{"role_assignment": {"next": 2}}`,
			expect: EpochState{
				RoleAssignmentEpochState: RoleAssignmentEpochState{
					Next: Layer2,
				},
				flag: roleAssignment,
			},
			err: false,
		},
		{
			name:  "Unknown string",
			input: `"unknown_state"`,
			err:   true,
		},
		{
			name:  "Unknown object",
			input: `{"unknown": {}}`,
			err:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var es EpochState
			err := es.UnmarshalText([]byte(tc.input))
			if tc.err {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				// Если ошибка ожидается, можно завершить тест
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			// Проверка для строковых состояний
			if tc.expect.IsInProgress() {
				if es.InProgressEpochState != tc.expect.InProgressEpochState {
					t.Errorf("Expected in_progress %q (InProgressEpochState = true), got %q", tc.expect.InProgressEpochState, es.InProgressEpochState)
				}
			}
			if tc.expect.IsReconcilingEvents() {
				if es.ReconcilingEventsEpochState != tc.expect.ReconcilingEventsEpochState {
					t.Errorf("Expected reconciling_events %q (ReconcilingEventsEpochState = true), got %q", tc.expect.ReconcilingEventsEpochState, es.ReconcilingEventsEpochState)
				}
			}
			// Проверка для объекта rewarding
			if tc.expect.IsRewarding() {
				if !reflect.DeepEqual(es.RewardingEpochState, tc.expect.RewardingEpochState) {
					t.Errorf("Expected rewarding %+v, got %+v", tc.expect.RewardingEpochState, es.RewardingEpochState)
				}
			}
			// Проверка для объекта role_assignment
			if tc.expect.IsRoleAssignment() {
				if !reflect.DeepEqual(es.RoleAssignmentEpochState, tc.expect.RoleAssignmentEpochState) {
					t.Errorf("Expected role_assignment %+v, got %+v", tc.expect.RoleAssignmentEpochState, es.RoleAssignmentEpochState)
				}
			}
		})
	}
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
			var dt OffsetDateTime
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
