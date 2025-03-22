package task

import (
	"fmt"
	"testing"
)

func TestNewTaskStatus(t *testing.T) {
	tests := []struct {
		value       int
		expected    int
		expectedErr error
	}{
		{value: 0, expected: 0, expectedErr: nil},
		{value: 1, expected: 1, expectedErr: nil},
		{value: 2, expected: 2, expectedErr: nil},
		{value: -1, expected: 0, expectedErr: fmt.Errorf("invalid TaskStatus value: %d", -1)},
		{value: 3, expected: 0, expectedErr: fmt.Errorf("invalid TaskStatus value: %d", 3)},
	}

	for _, test := range tests {
		taskStatus, err := NewTaskStatus(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTaskStatus(%d) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTaskStatus(%d) error = %v, want nil", test.value, err)
			}
			if taskStatus.Value() != test.expected {
				t.Errorf("NewTaskStatus(%d) value = %v, want %v", test.value, taskStatus.Value(), test.expected)
			}
		}
	}
}
