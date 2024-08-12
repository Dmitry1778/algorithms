package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	value := 7
	found := searchBST(root, value) // found = true
	fmt.Println(found)
}

func searchBST(root *TreeNode, value int) bool {
	if root == nil {
		return false
	}
	if root.Val == value {
		return true
	}
	if value < root.Val {
		return searchBST(root.Left, value)
	} else {
		return searchBST(root.Right, value)
	}
}
