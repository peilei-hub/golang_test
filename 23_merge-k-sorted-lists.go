package main

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeHelper(lists, 0, len(lists)-1)
}

func mergeHelper(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	listNode1 := mergeHelper(lists, left, mid)
	listNode2 := mergeHelper(lists, mid+1, right)

	return merge(listNode1, listNode2)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
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
