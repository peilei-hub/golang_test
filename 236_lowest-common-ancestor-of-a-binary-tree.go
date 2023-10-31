package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 把lowestCommonAncestor()当成是一个纯粹的二叉树后序遍历找节点函数(先忘记它是一个查找最近公共祖先的函数)，
// 目的是找到p或者q节点，然后根据返回结果来判断最近公共祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	// 从 root 找 p 或者 q

	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

func main() {
	r := &TreeNode{
		Val: 4,
	}

	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: r,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	tree := &TreeNode{
		Val:   3,
		Left:  p,
		Right: q,
	}

	ancestor := lowestCommonAncestor(tree, p, r)
	fmt.Println(ancestor)
}

var _236Res *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	findPorQ(root, p, q)
	return _236Res
}

func findPorQ(node, p, q *TreeNode) bool {
	if node == nil {
		return false
	}

	leftHasPorQ := findPorQ(node.Left, p, q)
	rightHasPorQ := findPorQ(node.Right, p, q)

	// 自下往上找的，所以一旦找到就是最小的公共祖先
	if leftHasPorQ && rightHasPorQ { //  如果这个节点的左右有 p 和 q，此节点就是公共祖先
		_236Res = node
	}
	if leftHasPorQ && (node == p || node == q) { // 如果这个节点左子节点有 p or q，且此节点是 p 或者 q，此节点就是公共祖先
		_236Res = node
	}
	if rightHasPorQ && (node == p || node == q) { // 如果这个节点右子节点有 p or q，且此节点是 p 或者 q，此节点就是公共祖先
		_236Res = node
	}

	return leftHasPorQ || rightHasPorQ || (node == p || node == q)
}
