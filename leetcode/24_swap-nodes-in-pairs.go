package main

// https://leetcode.cn/problems/swap-nodes-in-pairs/

func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{}

	result.Next = head

	pre := result
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := pre.Next.Next

		a.Next = b.Next
		pre.Next = b
		b.Next = a

		pre = pre.Next.Next
	}

	return result.Next
}
