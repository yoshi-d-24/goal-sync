package goal

import (
	"fmt"
	"testing"
)

func TestNewTitle(t *testing.T) {
	tests := []struct {
		value       string
		expected    string
		expectedErr error
	}{
		{value: "有効なタイトル", expected: "有効なタイトル", expectedErr: nil},
		{value: "", expected: "", expectedErr: fmt.Errorf("title must be at least %d characters long", MIN_LENGTH)},
		{value: "a", expected: "a", expectedErr: nil},
		{value: "a" + string(make([]rune, 50)), expected: "", expectedErr: fmt.Errorf("title must be no more than %d characters long", MAX_LENGTH)},
	}

	for _, test := range tests {
		title, err := NewTitle(test.value)
		if test.expectedErr != nil {
			if err == nil || err.Error() != test.expectedErr.Error() {
				t.Errorf("NewTitle(%s) error = %v, want %v", test.value, err, test.expectedErr)
			}
		} else {
			if err != nil {
				t.Errorf("NewTitle(%s) error = %v, want nil", test.value, err)
			}
			if title.Value() != test.expected {
				t.Errorf("NewTitle(%s) value = %v, want %v", test.value, title.Value(), test.expected)
			}
		}
	}
}
