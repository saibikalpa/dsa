package easy

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T){
	tests := []struct{
		str string
		want int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	funcs := []func(string) int{
		lengthOfLastWord,
		lengthOfLastWordOne,
	}
	for _, fn := range funcs{

		for _, test := range tests{
			t.Run(test.str, func(t *testing.T){
				actual := fn(test.str)
				if actual != test.want{
					t.Errorf("wanted %v but got %v",test.want, actual)
				}
			})
		}
	}
}
