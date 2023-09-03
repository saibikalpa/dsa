package trees

import "errors"

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
	size int
}

func (bt *BinaryTree) GetRoot() *BinaryTreeNode {
	return bt.root
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}
func (bt *BinaryTree) InOrderTraversal(tree *BinaryTree) []int {

	return nil
}
func (bt *BinaryTree) Insert(data int) error {
	if bt == nil {
		return errors.New("binary tree is nil")
	}
	bt.root = insert(bt.root, data)
	bt.size++
	return nil
}
func insert(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	}
	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}
func IdenticalTrees(t1 *BinaryTree, t2 *BinaryTree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return identicalTrees(t1.root, t2.root)
	}
	return false
}
func identicalTrees(r1 *BinaryTreeNode, r2 *BinaryTreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 != nil && r2 != nil {
		return r1.data == r2.data &&
			identicalTrees(r1.left, r2.left) &&
			identicalTrees(r1.right, r2.right)
	}
	return false
}
func (bt *BinaryTree) Height() int {
	if bt == nil {
		return 0
	}
	return height(bt.root)
}
func height(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	lh := height(root.left)
	rh := height(root.right)
	if lh >= rh {
		return lh + 1
	} else {
		return rh + 1
	}
}
func (bt *BinaryTree) LevelOrderTraversal() []int {
	if bt != nil {
		return levelOrderTraversal(bt.root)
	}
	return []int{}
}
func levelOrderTraversal(root *BinaryTreeNode) []int {
	h := height(root)
	res := make([]int, 0)
	for i := 1; i <= h; i++ {
		res = append(res, currentLevel(root, i)...)
	}
	return res
}
func (bt *BinaryTree) CurrentLevel(level int) []int {
	if bt != nil {
		return currentLevel(bt.root, level)
	}
	return []int{}
}

func currentLevel(t *BinaryTreeNode, l int) []int {
	res := make([]int, 0)
	if l == 1 {
		return []int{t.data}
	}
	a := currentLevel(t.left, l-1)
	b := currentLevel(t.right, l-1)
	res = append(res, a...)
	res = append(res, b...)
	return res
}
