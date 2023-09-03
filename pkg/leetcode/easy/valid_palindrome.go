package easy

import (
	"unicode"
)

func isPalindrome(s string) bool {
	str := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			str = append(str, unicode.ToLower(ch))
		} else if unicode.IsDigit(ch) {
			str = append(str, ch)
		}
	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
