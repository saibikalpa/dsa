package easy

import (
	"fmt"
	"strings"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	funcs := []func([]int) int{
		maxProfit,
		maxProfitOne}
	for _, fn := range funcs {
		for _, test := range tests {
			name := strings.Replace(fmt.Sprint(test.prices), " ", ",", -1)
			t.Run(name, func(t *testing.T) {
				actual := fn(test.prices)
				if actual != test.expected {
					t.Errorf("expected %v but got %v", test.expected, actual)
				}
			})
		}
	}

}
