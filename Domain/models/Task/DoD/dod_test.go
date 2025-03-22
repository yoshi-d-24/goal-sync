package task

import (
	"fmt"
	"testing"
)

func TestNewDoD(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "有効なDoD", expected: "有効なDoD", expectedErr: nil},
		{value: "", expected: "", expectedErr: fmt.Errorf("DoD must be at least %d characters long", MIN_LENGTH)},
		{value: "a", expected: "a", expectedErr: nil},
		{value: string(make([]rune, 301)), expected: "", expectedErr: fmt.Errorf("DoD must be no more than %d characters long", MAX_LENGTH)},
	}

	for _, test := range tests {
		dod, err := NewDoD(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewDoD(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewDoD(%s) error = %v, want nil", test.value, err)
			}
			if dod.Value() != test.expected {
				t.Errorf("NewDoD(%s) value = %v, want %v", test.value, dod.Value(), test.expected)
			}
		}
	}
}
