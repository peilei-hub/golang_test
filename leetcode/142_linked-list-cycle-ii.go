package main

// https://leetcode.cn/problems/linked-list-cycle-ii/description/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			result := head

			for result != slow {
				result = result.Next
				slow = slow.Next
			}
			return result
		}
	}

	return nil
}
