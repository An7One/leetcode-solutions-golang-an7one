// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
//
// Time Complexity:		O(N)
// Space Complexity:	O(H)
//
// Reference:
// 	https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/765931/Go-golang-recursive-and-iterative-solutions
package lc0589

import "github.com/an7one/leetcode-solutions-golang-an7one/src/leetcode/util/ds"

type Node = ds.NarryTreeNode

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack, ans := []*Node{root}, []int{}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, pop.Val)
		for i := len(pop.Children) - 1; i >= 0; i-- {
			stack = append(stack, pop.Children[i])
		}
	}

	return ans
}
