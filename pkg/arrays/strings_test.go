package arrays

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"aba", true},
		{"rotator", true},
		{"bs", false},
		{"madam", true},
		{"lilly", false},
	}

	for _, test := range tests {
		actual := IsPalindrome(test.word)
		if actual != test.expected {
			t.Errorf("Error encountered at IsPalindrome: expected %v but got %v for %v", test.expected, actual, test.word)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"apple", "elppa"},
		{"red", "der"},
		{"deer", "reed"},
		{"gang", "gnag"},
		{"home", "emoh"},
	}

	for _, test := range tests {
		t.Run(test.word, func(t *testing.T) {
			actual := Reverse(test.word)
			if actual != test.expected {
				t.Errorf("Error at Reverse: expected %v but got %v", test.expected, actual)
			}
		})
	}
}
