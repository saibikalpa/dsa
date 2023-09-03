package easy

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	a := inorderTraversal(root.Left)
	b := inorderTraversal(root.Right)
	res := make([]int, 0)
	res = append(res, a...)
	res = append(res, root.Val)
	res = append(res, b...)
	return res
}
