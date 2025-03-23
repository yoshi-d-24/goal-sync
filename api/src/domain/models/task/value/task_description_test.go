package value

import (
	"fmt"
	"testing"
)

func TestNewTaskDescription(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "有効なタスク詳細", expected: "有効なタスク詳細", expectedErr: nil},
		{value: "a", expected: "a", expectedErr: nil},
		{value: string(make([]rune, 201)), expected: "", expectedErr: fmt.Errorf("TaskDescription must be no more than %d characters long", TASK_DESCRIPTION_MAX_LENGTH)},
	}

	for _, test := range tests {
		taskDescription, err := NewTaskDescription(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTaskDescription(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTaskDescription(%s) error = %v, want nil", test.value, err)
			}
			if taskDescription.Value() != test.expected {
				t.Errorf("NewTaskDescription(%s) value = %v, want %v", test.value, taskDescription.Value(), test.expected)
			}
		}
	}
}
