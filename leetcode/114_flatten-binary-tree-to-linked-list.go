package main

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/

func flatten(root *leetcode.TreeNode) {
	if root == nil {
		return
	}

	// 利用定义，把左右子树拉平
	flatten(root.Left)
	flatten(root.Right)

	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right

	// 2. 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3. 将原先右子树拼到当前右子树末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
