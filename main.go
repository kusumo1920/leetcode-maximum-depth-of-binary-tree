package main

import "fmt"

func main() {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	output := maxDepthFromBottomUp(input)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthFromBottomUp(root *TreeNode) int {
	var recursiveFn func(*TreeNode, int) int
	recursiveFn = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth - 1
		}
		leftDepth := recursiveFn(node.Left, depth+1)
		rightDepth := recursiveFn(node.Right, depth+1)
		return max(leftDepth, rightDepth)
	}

	return recursiveFn(root, 1)
}
