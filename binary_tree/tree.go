package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tree = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

func main() {
	var res1 []int
	preOrder1(tree, &res1)
	fmt.Println(res1)

	var res2 []int
	preOrder2(tree, &res2)
	fmt.Println(res2)

	var res3 []int
	midOrder1(tree, &res3)
	fmt.Println(res3)

	var res4 []int
	midOrder2(tree, &res4)
	fmt.Println(res4)

	var res5 []int
	behindOrder1(tree, &res5)
	fmt.Println(res5)

	var res6 []int
	behindOrder2(tree, &res6)
	fmt.Println(res6)

	var res7 []int
	widthOrder(tree, &res7)
	fmt.Println(res7)
}

func preOrder1(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	*res = append(*res, head.Val)

	preOrder1(head.Left, res)
	preOrder1(head.Right, res)
}

func preOrder2(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, head)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*res = append(*res, node.Val) // print root

		if node.Right != nil {
			stack = append(stack, node.Right) // push right
		}
		if node.Left != nil {
			stack = append(stack, node.Left) // push left
		}
	}
}

func midOrder1(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	midOrder1(head.Left, res)
	*res = append(*res, head.Val)
	midOrder1(head.Right, res)
}

func midOrder2(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	temp := head
	for temp != nil { // 从head开始将左全部入栈
		stack = append(stack, temp)
		temp = temp.Left
	}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		*res = append(*res, cur.Val)

		temp2 := cur.Right // 从 cur.Right 开始将左全部入栈
		for temp2 != nil {
			stack = append(stack, temp2)
			temp2 = temp2.Left
		}
	}
}

func behindOrder1(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	behindOrder1(head.Left, res)
	behindOrder1(head.Right, res)
	*res = append(*res, head.Val)
}

func behindOrder2(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)

	stack1 = append(stack1, head)
	for len(stack1) != 0 {
		cur := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		stack2 = append(stack2, cur)
		if cur.Left != nil {
			stack1 = append(stack1, cur.Left)
		}
		if cur.Right != nil {
			stack1 = append(stack1, cur.Right)
		}
	}

	for len(stack2) != 0 {
		cur := stack2[len(stack2)-1]
		*res = append(*res, cur.Val)
		stack2 = stack2[:len(stack2)-1]
	}
}

func widthOrder(head *TreeNode, res *[]int) {
	if head == nil {
		return
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, head)

	for len(queue) != 0 {
		cur := queue[0]
		*res = append(*res, cur.Val)
		queue = queue[1:]

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}
