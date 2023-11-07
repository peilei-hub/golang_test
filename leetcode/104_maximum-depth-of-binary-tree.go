package main

func maxDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	return max104(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}
