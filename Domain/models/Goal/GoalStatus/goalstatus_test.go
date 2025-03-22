package goal

import (
	"fmt"
	"testing"
)

func TestNewGoalStatus(t *testing.T) {
	tests := []struct {
		value       int
		expected    GoalStatusEnum
		expectedErr error
	}{
		{value: 0, expected: Incomplete, expectedErr: nil},
		{value: 1, expected: Complete, expectedErr: nil},
		{value: -1, expected: 0, expectedErr: fmt.Errorf("invalid GoalStatusEnum value: %d", -1)},
		{value: 2, expected: 0, expectedErr: fmt.Errorf("invalid GoalStatusEnum value: %d", 2)},
	}

	for _, test := range tests {
		goalStatus, err := NewGoalStatus(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewGoalStatus(%d) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewGoalStatus(%d) error = %v, want nil", test.value, err)
			}
			if goalStatus.Value() != test.expected {
				t.Errorf("NewGoalStatus(%d) value = %v, want %v", test.value, goalStatus.Value(), test.expected)
			}
		}
	}
}
