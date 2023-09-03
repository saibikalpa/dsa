package arrays

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLinearSearch(t *testing.T) {

	tests := []struct {
		array    []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 78, 9}, 78, 3},
		{[]int{13, 34, 25, 78, 29}, 34, 1},
		{[]int{10, 30, 50, 78, 90}, 78, 3},
		{[]int{21, 23, 25, 28, 29}, 22, -1},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := LinearSearch(test.array, test.target)
			if actual != test.expected {
				t.Errorf("error occured in Linear search: expected %v but got %v", test.expected, actual)
			}
		})
	}
}

func TestBinarySearch(t *testing.T){
	tests:=[]struct{
		array []int
		target int
		expected int
	}{
		{[]int{1, 3, 5, 78, 9}, 78, 3},
		{[]int{13, 34, 35, 78, 29}, 34, 1},
		{[]int{10, 30, 50, 78, 90}, 78, 3},
		{[]int{21, 23, 25, 28, 29}, 22, -1},
	}
	for _, test:=range tests{
		t.Run(fmt.Sprint(test.array), func(t *testing.T){
			actual := BinarySearch(test.array, test.target)
			if actual != test.expected{
				t.Errorf("error occuured in binary search: expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
