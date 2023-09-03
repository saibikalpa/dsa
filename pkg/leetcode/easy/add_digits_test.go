package easy

import (
	"testing"
)

func Test_addDigits(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{38, 2},
		{0, 0},
		{9, 9},
		{35, 8},
	}
	funcs := []func(int) int{
		addDigits,
		addDigitsOne,
		addDigitsTwo,
	}
	for _, fn := range funcs {
		for _, test := range tests {
			actual := fn(test.number)
			if actual != test.want {
				t.Errorf("wanted %v, but got %v", test.want, actual)
			}
		}
	}
}
