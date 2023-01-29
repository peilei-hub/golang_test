package main

// https://leetcode.cn/problems/merge-two-sorted-lists/

/**
 * Definition for singly-linked list.
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list1Temp := list1
	list2Temp := list2

	result := &ListNode{}
	var cur = result
	for list1Temp != nil && list2Temp != nil {
		if list1Temp.Val < list2Temp.Val {
			cur.Next = list1Temp
			list1Temp = list1Temp.Next
		} else {
			cur.Next = list2Temp
			list2Temp = list2Temp.Next
		}
		cur = cur.Next
	}

	if list1Temp != nil {
		cur.Next = list1Temp
	}

	if list2Temp != nil {
		cur.Next = list2Temp
	}

	return result.Next
}
