package easy

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {

	testcase1 := func() *TreeNode {
		t := &TreeNode{1, nil, nil}
		t.Left = nil
		t.Right = &TreeNode{2, nil, nil}
		t.Right.Left = &TreeNode{3, nil, nil}
		return t
	}
	tests := []struct {
		name     string
		fn       func() *TreeNode
		expected []int
	}{
		{
			name:     "Test Case 1",
			fn:       testcase1,
			expected: []int{1, 3, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := inorderTraversal(test.fn())
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v but got %v", test.expected, actual)
			}
		})
	}

}
