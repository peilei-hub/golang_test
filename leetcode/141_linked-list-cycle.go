package main

// https://leetcode.cn/problems/linked-list-cycle/description/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
