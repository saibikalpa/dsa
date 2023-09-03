package easy

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			break
		}
		i--
	}
	return len(s) - i - 1
}

func lengthOfLastWordOne(s string) int {
	size := len(s)
	length := 0

	for i:=size -1;i>= 0; i--{
		if length == 0 && s[i]==' '{
			continue
		}else{
			if s[i] == ' '{
				break
			}
		}
		length++
	}
	return length
}
