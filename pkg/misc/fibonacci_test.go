package misc

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
		{7, 8},
	}
	funcs := []func(int) int{
		Fibonacci,
		FibonacciOne,
		FibonacciTwo,
	}
	for _, fn := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.n), func(t *testing.T) {
				actual := fn(test.n)
				if actual != test.want {
					t.Errorf("expected %v but got %v", test.want, actual)
				}
			})
		}
	}
}
