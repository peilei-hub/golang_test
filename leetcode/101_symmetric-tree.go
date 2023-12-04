package main

// https://leetcode.cn/problems/symmetric-tree/description/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return d101(root.Left, root.Right)
}

func d101(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return d101(left.Left, right.Right) && d101(left.Right, right.Left) // 左节点的左孩子和右节点的右孩子，以及左节点的右孩子和右节点的左孩子
}
