package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {

	testInt := []int{1, 2, 4, 6, 8, 9, 14, 15}
	targetSum := 29
	expected := true
	actual := TwoSumSorted(testInt, targetSum)
	if expected != actual {
		t.Errorf("Failed!\n Correct result is %v\n", expected)
	}
}

func Test_mergeSortedArrays(t *testing.T) {
	tests := []struct {
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			arr1:     []int{1, 3, 5, 7, 9},
			arr2:     []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arr1:     []int{1, 13, 15, 17, 19},
			arr2:     []int{2, 4, 16, 28},
			expected: []int{1, 2, 4, 13, 15, 16, 17, 19, 28},
		},
	}

	for _, test := range tests {
		actual := MergeSortedArrays(test.arr1, test.arr2)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("merge failed: expected %v, but got %v", test.expected, actual)
		}
	}
}
