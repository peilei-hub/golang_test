package main

// 把lowestCommonAncestor()当成是一个纯粹的二叉树后序遍历找节点函数(先忘记它是一个查找最近公共祖先的函数)，
// 目的是找到p或者q节点，然后根据返回结果来判断最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 从 root 找 p 或者 q

	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

func findPorQ(node, p, q *TreeNode) *TreeNode {
	if node == nil || node == p || node == q {
		return node
	}

	left := findPorQ(node.Left, p, q)
	right := findPorQ(node.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return node
}
