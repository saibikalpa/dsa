package easy

import "sort"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charS := []rune(s)

	charT := []rune(t)

	sort.SliceStable(charS, func(i, j int) bool {
		return charS[i] < charS[j]
	})
	sort.SliceStable(charT, func(i, j int) bool {
		return charT[i] < charT[j]
	})
	for i := 0; i < len(charS); i++ {
		if charS[i] != charT[i] {
			return false
		}
	}
	return true
}

func isAnagramOne(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make([]int, 26)

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, d := range freq {
		if d != 0 {
			return false
		}
	}
	return true
}
