package goal

import (
	"fmt"
	"testing"
)

func TestNewGoalDescription(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "有効な詳細", expected: "有効な詳細", expectedErr: nil},
		{value: "", expected: "", expectedErr: fmt.Errorf("goal description must be at least %d characters long", MIN_LENGTH)},
		{value: "a", expected: "a", expectedErr: nil},
		{value: "a" + string(make([]rune, 500)), expected: "", expectedErr: fmt.Errorf("goal description must be no more than %d characters long", MAX_LENGTH)},
	}

	for _, test := range tests {
		goalDescription, err := NewGoalDescription(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewGoalDescription(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewGoalDescription(%s) error = %v, want nil", test.value, err)
			}
			if goalDescription.Value() != test.expected {
				t.Errorf("NewGoalDescription(%s) value = %v, want %v", test.value, goalDescription.Value(), test.expected)
			}
		}
	}
}
