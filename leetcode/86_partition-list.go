package main

// https://leetcode.cn/problems/partition-list/

func partition(head *ListNode, x int) *ListNode {
	// 拆成两个链表再合并

	list1Node := &ListNode{}
	list2Node := &ListNode{}

	result := list1Node
	list2Pre := list2Node

	node := head

	for node != nil {
		if node.Val < x {
			list1Node.Next = node
			node = node.Next

			list1Node = list1Node.Next
			list1Node.Next = nil
		} else {
			list2Node.Next = node
			node = node.Next

			list2Node = list2Node.Next
			list2Node.Next = nil
		}
	}

	list1Node.Next = list2Pre.Next

	return result.Next
}
