// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
//
// Time Complexity:		O(N)
// Space Complexity:	O(H)
//
// Reference:
//	https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/765931/Go-golang-recursive-and-iterative-solutions
package lc0589

import "github.com/an7one/leetcode-solutions-golang-an7one/src/leetcode/util/ds"

type Node = ds.NarryTreeNode

func preorder(root *Node) []int {
	ans := []int{}

	dfs(root, &ans)

	return ans
}

// preorder recursive
func dfs(node *Node, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)

	for _, child := range node.Children {
		dfs(child, res)
	}
}
