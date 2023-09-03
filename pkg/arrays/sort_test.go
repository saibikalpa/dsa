package arrays

import (
	"fmt"
	"strings"
	"testing"
	"reflect"
	"runtime"
)

func TestSorting(t *testing.T){
	tests := []struct{
		arr []int
		want []int
	}{
		{
			arr: []int{5,4,3,2,1},
			want: []int{1,2,3,4,5},
		},
		{
			arr:[]int{2, 0, 7, 1},
			want:[]int{0, 1, 2, 7},
		},
		{
			arr:[]int{2,-9,8,10,-34},
			want:[]int{-34, -9, 2, 8, 10},
		},
		{
			arr:[]int{100,200, 99, 87, 54},
			want:[]int{54, 87, 99, 100, 200},
		},
		{
			arr:[]int{20, 10, -5, -6},
			want:[]int{-6, -5, 10, 20},
		},
		{
			arr:[]int{0},
			want:[]int{0},
		},
	}

	funcs := []func([]int){
		BubbleSort,
		SelectionSort,
		InsertionSort,
	}
	for _, fn:= range funcs{
		for _, test:= range tests{
			t.Run(strings.Replace(fmt.Sprint(test.arr), " ", ",", -1), func(t *testing.T){
				actual := make([]int, len(test.arr))
				copy(actual, test.arr)
				fn(actual)
				if !reflect.DeepEqual(actual , test.want){
					t.Errorf("%v failed: expected %v but got %v",GetFunctionName(fn), test.want, test.arr)
				}
			})
		}
	}
}
func GetFunctionName(f interface{}) string{
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
