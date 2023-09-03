package easy

import (
	"testing"
	"fmt"
)

func Test_isUgly(t *testing.T){
	tests:= []struct{

		number int
		want bool
	}{
		{6, true},
		{0, false},
		{2, true},
		{14, false},
		{-2, false},
	}
	for _,test := range tests{
		t.Run(fmt.Sprint(test.number), func(t *testing.T){
			actual := isUgly(test.number)
			if actual != test.want{
				t.Errorf("expected %v but got %v",test.want, actual)
			}
		})
	}
}
