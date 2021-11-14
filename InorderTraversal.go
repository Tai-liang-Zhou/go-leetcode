<<<<<<< HEAD
package main

import (
	"fmt"
)

// Definition for a binary tree node.
=======
// package main

// import (
// 	"fmt"
// )

// // Definition for a binary tree node.
>>>>>>> 42311c82b44eac0505257d83ef512e19c7dd9976
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

<<<<<<< HEAD
//
type BTree struct {
	Root *TreeNode
}

func InoderTraveral(root *TreeNode) []int {
	var xs []int
	if root != nil {
		xs = append(xs, InoderTraveral(root.Left)...)
		xs = append(xs, root.Val)
		xs = append(xs, InoderTraveral(root.Right)...)
		fmt.Println(xs)
	}
	return xs
}

func InertTreeNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{value, nil, nil}
	}

	if root.Val > value {
		root.Left = InertTreeNode(root.Left, value)
	} else {
		root.Right = InertTreeNode(root.Right, value)
	}

	return root
}
=======
// //
// type BTree struct {
// 	Root *TreeNode
// }

// func InoderTraveral(root *TreeNode) []int {
// 	var xs []int
// 	if root != nil {
// 		xs = append(xs, InoderTraveral(root.Left)...)
// 		xs = append(xs, root.Val)
// 		xs = append(xs, InoderTraveral(root.Right)...)
// 		fmt.Println(xs)
// 	}
// 	return xs
// }

// func InertTreeNode(root *TreeNode, value int) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{value, nil, nil}
// 	}

// 	if root.Val > value {
// 		root.Left = InertTreeNode(root.Left, value)
// 	} else {
// 		root.Right = InertTreeNode(root.Right, value)
// 	}

// 	return root
// }
>>>>>>> 42311c82b44eac0505257d83ef512e19c7dd9976

// func main() {
// 	t := &BTree{nil}
// 	t.Root = InertTreeNode(t.Root, 2)
// 	t.Root = InertTreeNode(t.Root, 3)
// 	t.Root = InertTreeNode(t.Root, 4)

// 	xs := InoderTraveral(t.Root)
// 	fmt.Println(xs)
// 	fmt.Printf("root value :" + strconv.Itoa(t.Root.Val))

// }
