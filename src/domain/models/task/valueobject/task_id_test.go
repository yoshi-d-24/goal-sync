package valueobject

import (
	"fmt"
	"testing"
)

func TestNewTaskId(t *testing.T) {
	tests := []struct {
		value       int
		expected    int
		expectedErr error
	}{
		{value: 1, expected: 1, expectedErr: nil},
		{value: 100, expected: 100, expectedErr: nil},
		{value: 0, expected: 0, expectedErr: fmt.Errorf("TaskId must be greater than 0")},
		{value: -1, expected: -1, expectedErr: fmt.Errorf("TaskId must be greater than 0")},
	}

	for _, test := range tests {
		taskId, err := NewTaskId(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTaskId(%d) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTaskId(%d) error = %v, want nil", test.value, err)
			}
			if taskId.Value() != test.expected {
				t.Errorf("NewTaskId(%d) value = %v, want %v", test.value, taskId.Value(), test.expected)
			}
		}
	}
}
