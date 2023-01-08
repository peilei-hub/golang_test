package main

// https://leetcode.cn/problems/add-two-numbers/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	result := &ListNode{}
	temp := result
	for l1 != nil || l2 != nil {
		var l1Val = 0
		var l2Val = 0

		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		sum := l1Val + l2Val + carry

		val := sum % 10
		carry = sum / 10

		temp.Val = val

		if l1 != nil || l2 != nil {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}
	if carry > 0 {
		temp.Next = &ListNode{}
		temp.Next.Val = carry
	}

	return result
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	result := &ListNode{}
	temp := result
	for l1 != nil || l2 != nil {
		var l1Val = 0
		var l2Val = 0

		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		sum := l1Val + l2Val + carry

		val := sum % 10
		carry = sum / 10

		temp.Next = &ListNode{}
		temp = temp.Next
		temp.Val = val
	}
	if carry > 0 {
		temp.Next = &ListNode{}
		temp.Next.Val = carry
	}

	return result.Next
}
