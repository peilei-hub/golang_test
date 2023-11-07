package main

// https://leetcode.cn/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	newHead := reverseNode(head, tail) // 获得部分反转后的头结点，此时的tail为另外未反转的头
	head.Next = reverseKGroup(tail, k) // 此时head为部分反转的尾部，将尾部的next=递归的结果
	return newHead
}

func reverseNode(head, tail *ListNode) *ListNode { // 将从 [head,tail)的顺序翻转，并返回反转后的头结点，此时head为反转后的尾节点
	cur := head
	var curPre *ListNode

	for cur != tail {
		next := cur.Next
		cur.Next = curPre

		curPre = cur
		cur = next
	}

	return curPre
}
