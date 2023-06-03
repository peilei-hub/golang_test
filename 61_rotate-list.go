package main

// https://leetcode.cn/problems/rotate-list/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil { // 找到tail节点，以及链表的长度
		length++
		tail = tail.Next
	}

	k %= length
	if k == 0 {
		return head
	}

	cur := head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}

	kPre := head
	for cur.Next != nil {
		cur = cur.Next
		kPre = kPre.Next
	}

	kNode := kPre.Next
	kPre.Next = nil
	tail.Next = head

	return kNode
}
