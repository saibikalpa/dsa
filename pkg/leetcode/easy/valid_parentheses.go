package easy

import "github.com/saibikalpa/dsa/pkg/stacks"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	stk := stacks.NewStack()
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stk.Push(ch)
		} else {
			v, _ := stk.Pop()
			if ch == ')' && v != '(' || ch == ']' && v != '[' || ch == '}' && v != '{' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}

func isValidOne(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stk := stacks.NewStack()
	for _, ch := range s {
		if _, keyFound := pairs[ch]; keyFound {
			stk.Push(ch)
		} else if t, e := stk.Pop(); e != nil || ch != pairs[t.(rune)] {
			return false
		}
	}
	return stk.IsEmpty()
}

func isValidTwo(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stk := []rune{}
	for _, ch := range s {
		if _, keyFound := pairs[ch]; keyFound {
			stk = append(stk, ch)
		} else if len(stk) == 0 || ch != pairs[stk[len(stk)-1]] {
			return false
		} else {
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
