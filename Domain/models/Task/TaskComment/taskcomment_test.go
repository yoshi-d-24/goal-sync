package task

import (
	"fmt"
	"testing"
)

func TestNewTaskComment(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "有効なコメント", expected: "有効なコメント", expectedErr: nil},
		{value: "", expected: "", expectedErr: fmt.Errorf("TaskComment must be at least %d characters long", MIN_LENGTH)},
		{value: "a", expected: "a", expectedErr: nil},
		{value: string(make([]rune, 1001)), expected: "", expectedErr: fmt.Errorf("TaskComment must be no more than %d characters long", MAX_LENGTH)},
	}

	for _, test := range tests {
		taskComment, err := NewTaskComment(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTaskComment(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTaskComment(%s) error = %v, want nil", test.value, err)
			}
			if taskComment.Value() != test.expected {
				t.Errorf("NewTaskComment(%s) value = %v, want %v", test.value, taskComment.Value(), test.expected)
			}
		}
	}
}
