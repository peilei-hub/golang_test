package main

import "fmt"

// https://leetcode.cn/problems/reverse-linked-list-ii/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	leftNode := head
	rightNode := head

	pre := &ListNode{}
	pre.Next = head
	res := pre
	for i := 0; i < left-1; i++ { // pre为left的前驱节点
		if leftNode == nil {
			return head
		}
		pre = pre.Next
		leftNode = leftNode.Next
	}

	for i := 0; i < right; i++ {
		if rightNode == nil {
			return head
		}
		rightNode = rightNode.Next // rightNode为right的后继节点
	}

	h := reverse92(leftNode, rightNode) // 返回后h为翻转部分的头，leftNode为翻转的尾部

	pre.Next = h
	leftNode.Next = rightNode

	return res.Next
}

func reverse92(head, tail *ListNode) (h *ListNode) {
	cur := head
	var preCur *ListNode

	for cur != tail { // 不包括尾部
		next := cur.Next

		cur.Next = preCur

		preCur = cur
		cur = next
	}

	return preCur
}

func main() {
	list1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}

	between := reverseBetween(list1, 1, 2)
	fmt.Println(between)
}
