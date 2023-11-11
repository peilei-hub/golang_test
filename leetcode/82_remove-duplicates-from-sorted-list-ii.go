package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	result := &ListNode{
		Next: head,
	}

	pre := result

	for pre.Next != nil && pre.Next.Next != nil { // 下一个和下下个是否相同
		if pre.Next.Val == pre.Next.Next.Val { // 相同
			val := pre.Next.Val

			temp := pre.Next.Next.Next
			for temp != nil && temp.Val == val {
				temp = temp.Next
			}

			pre.Next = temp // 将相同的遍历过去
		} else {
			pre = pre.Next // 不同则后移一位
		}
	}

	return result.Next
}
