package main

func invertTree(root *leetcode.TreeNode) *leetcode.TreeNode {
	traverse226(root)
	return root
}

func invertTree2(root *leetcode.TreeNode) *leetcode.TreeNode {
	if root == nil {
		return root
	}

	left := invertTree2(root.Left)
	right := invertTree2(root.Right)
	root.Left = right
	root.Right = left

	return root
}

func traverse226(root *leetcode.TreeNode) {
	if root == nil {
		return
	}

	traverse226(root.Left)
	traverse226(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
