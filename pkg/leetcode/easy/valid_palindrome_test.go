package easy

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			actual := isPalindrome(test.str)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}
