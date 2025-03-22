package goal

import (
	"fmt"
	"testing"
)

func TestNewGoalStatus(t *testing.T) {
	tests := []struct {
		value       int
		expected    int
		expectedErr error
	}{
		{value: 1, expected: 1, expectedErr: nil},
		{value: 5, expected: 5, expectedErr: nil},
		{value: 10, expected: 10, expectedErr: nil},
		{value: 0, expected: 0, expectedErr: fmt.Errorf("importance must be between %d and %d", MIN, MAX)},
		{value: 11, expected: 0, expectedErr: fmt.Errorf("importance must be between %d and %d", MIN, MAX)},
	}

	for _, test := range tests {
		importance, err := NewGoalStatus(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewGoalStatus(%d) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewGoalStatus(%d) error = %v, want nil", test.value, err)
			}
			if importance.Value() != test.expected {
				t.Errorf("NewGoalStatus(%d) value = %v, want %v", test.value, importance.Value(), test.expected)
			}
		}
	}
}
