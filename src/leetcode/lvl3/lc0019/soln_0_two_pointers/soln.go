package lc0019

import "github.com/an7one/leetcode-solutions-golang-an7one/src/leetcode/util/ds"

type ListNode = ds.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}

	hi := dummy
	lo := dummy

	for i := 1; i <= n+1; i++ {
		hi = hi.Next
	}

	for hi != nil {
		hi = hi.Next
		lo = lo.Next
	}

	lo.Next = lo.Next.Next

	return dummy.Next
}
