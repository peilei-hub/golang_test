package main

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	midTraversal(root, &result)
	return result
}

func midTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	midTraversal(node.Left, result)
	*result = append(*result, node.Val)
	midTraversal(node.Right, result)
}
