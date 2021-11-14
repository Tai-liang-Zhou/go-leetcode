package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSemmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return visite(root.Left, root.Right)
}

func visite(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return visite(root1.Right, root2.Left) && visite(root1.Left, root2.Right)
}

// func main() {

// }
