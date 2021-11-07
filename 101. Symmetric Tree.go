package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isSemmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return visite(root.left, root.right)
}

func visite(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.val != root2.val {
		return false
	}

	return visite(root1.right, root2.left) && visite(root1.left, root2.right)
}

func main() {

}
