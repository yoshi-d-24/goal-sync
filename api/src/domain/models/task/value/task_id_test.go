package value

import (
	"fmt"
	"testing"
)

func TestNewTaskId(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "50ac2aa3-ab64-4184-9112-d23221dc1832", expected: "50ac2aa3-ab64-4184-9112-d23221dc1832", expectedErr: nil},
		{value: "", expected: "", expectedErr: fmt.Errorf("TaskId must be uuid")},
		{value: "test", expected: "", expectedErr: fmt.Errorf("TaskId must be uuid")},
	}

	for _, test := range tests {
		taskId, err := NewTaskId(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTaskId(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTaskId(%s) error = %v, want nil", test.value, err)
			}
			if taskId.Value() != test.expected {
				t.Errorf("NewTaskId(%s) value = %v, want %v", test.value, taskId.Value(), test.expected)
			}
		}
	}
}
