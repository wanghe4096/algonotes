package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println("vim-go")
}
