package main

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// preNode防止删除head导致空指针问题
	pre := &ListNode{
		Next: head,
	}

	l := pre
	r := pre

	// 先走n步
	for i := 0; i < n; i++ {
		r = r.Next
	}

	// 将l定位到要删除的前一个节点
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return pre.Next
}
