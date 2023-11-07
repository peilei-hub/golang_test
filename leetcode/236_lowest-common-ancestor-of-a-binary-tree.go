package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res TreeNode
	findPorQ(root, p, q, &res)
	return &res
}

func findPorQ(node, p, q, res *TreeNode) bool {
	if node == nil {
		return false
	}

	leftHasPorQ := findPorQ(node.Left, p, q, res)
	rightHasPorQ := findPorQ(node.Right, p, q, res)

	// 自下往上找的，所以一旦找到就是最小的公共祖先
	if leftHasPorQ && rightHasPorQ { //  如果这个节点的左右有 p 和 q，此节点就是公共祖先
		*res = *node
	}
	if leftHasPorQ && (node == p || node == q) { // 如果这个节点左子节点有 p or q，且此节点是 p 或者 q，此节点就是公共祖先
		*res = *node
	}
	if rightHasPorQ && (node == p || node == q) { // 如果这个节点右子节点有 p or q，且此节点是 p 或者 q，此节点就是公共祖先
		*res = *node
	}

	return leftHasPorQ || rightHasPorQ || (node == p || node == q)
}
