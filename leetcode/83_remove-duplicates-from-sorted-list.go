package main

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val { // 当前val 和下一个val相等
			cur.Next = cur.Next.Next // 将下一个节点改成下下个节点
		} else { // cur.Val != cur.Next.Val
			cur = cur.Next
		}
	}

	return head
}
