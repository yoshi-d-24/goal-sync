package goal

import (
	"fmt"
	"testing"
)

func TestNewGoalId(t *testing.T) {
	tests := []struct {
		value       int
		expected    int
		expectedErr error
	}{
		{value: 1, expected: 1, expectedErr: nil},
		{value: 100, expected: 100, expectedErr: nil},
		{value: 0, expected: 0, expectedErr: fmt.Errorf("goal id should be greater than 0")},
		{value: -1, expected: -1, expectedErr: fmt.Errorf("goal id should be greater than 0")},
	}

	for _, test := range tests {
		goalId, err := NewGoalId(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewGoalId(%d) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewGoalId(%d) error = %v, want nil", test.value, err)
			}
			if goalId.Value() != test.expected {
				t.Errorf("NewGoalId(%d) value = %v, want %v", test.value, goalId.Value(), test.expected)
			}
		}
	}
}
