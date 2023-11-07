package main

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	traverse116(root.Left, root.Right)
	return root
}

func traverse116(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	traverse116(node1.Left, node1.Right)
	traverse116(node2.Left, node2.Right)
	traverse116(node1.Right, node2.Left)
}
