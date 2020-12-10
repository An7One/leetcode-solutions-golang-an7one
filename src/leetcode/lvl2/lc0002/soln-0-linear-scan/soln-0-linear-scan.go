package lc0002

import (
	"leetcode-solutions-golang-zea7ot/src/leetcode/util/ds"
)

// ListNode is the node for a singly-list node
type ListNode = ds.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	dummy := &ListNode{Val: -1, Next: nil}
	prev := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry
		prev.Next = &ListNode{Val: sum % 10, Next: nil}
		prev = prev.Next
		carry = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		prev.Next = &ListNode{Val: carry, Next: nil}
	}

	return dummy.Next
}
