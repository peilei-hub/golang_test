package main

// https://leetcode.cn/problems/rotate-list/

// 1. 注意链表为空，k=0
// 2. 注意k大于链表长度，需要取模

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

	kPre := head // 找到倒数第k个结点的前驱结点
	for cur.Next != nil {
		cur = cur.Next
		kPre = kPre.Next
	}

	kNode := kPre.Next // 倒数第k个结点，为新的头结点
	kPre.Next = nil
	tail.Next = head

	return kNode
}
