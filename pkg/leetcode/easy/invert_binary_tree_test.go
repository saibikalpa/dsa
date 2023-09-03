package easy

import (
	"testing"
)

func Test_invertTree(t *testing.T) {
	testInput1 := func() *TreeNode {
		// [4,2,7,1,3,6,9]
		t := &TreeNode{4, nil, nil}
		t.Left = &TreeNode{2, nil, nil}
		t.Right = &TreeNode{7, nil, nil}
		t.Left.Left = &TreeNode{1, nil, nil}
		t.Left.Right = &TreeNode{3, nil, nil}
		t.Right.Left = &TreeNode{6, nil, nil}
		t.Right.Right = &TreeNode{9, nil, nil}
		return t
	}
	testOutput1 := func() *TreeNode {
		// [4,7,2,9,6,3,1]
		t := &TreeNode{4, nil, nil}
		t.Left = &TreeNode{7, nil, nil}
		t.Right = &TreeNode{2, nil, nil}
		t.Left.Left = &TreeNode{9, nil, nil}
		t.Left.Right = &TreeNode{6, nil, nil}
		t.Right.Left = &TreeNode{3, nil, nil}
		t.Right.Right = &TreeNode{1, nil, nil}
		return t
	}
	tests := []struct {
		name  string
		input func() *TreeNode
		want  func() *TreeNode
	}{
		{
			name:  "Test 1",
			input: testInput1,
			want:  testOutput1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !identicalTrees(invertTree(test.input()), test.want()) {
				t.Errorf("trees do not match, error occured while inverting tree")
			}
		})
	}
}
func identicalTrees(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 != nil && r2 != nil {
		return r1.Val == r2.Val &&
			identicalTrees(r1.Left, r2.Left) &&
			identicalTrees(r1.Right, r2.Right)
	}
	return false
}
