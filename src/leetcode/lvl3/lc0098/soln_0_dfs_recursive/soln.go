// @author: Leon
// https://leetcode.com/problems/validate-binary-search-tree/
//
// Time Complexity:		O(N)
// Space Complexity:	O(H)
package lc0098

import (
	"github.com/an7one/leetcode-solutions-golang-an7one/src/leetcode/util/ds"
)

type TreeNode = ds.TreeNode

func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

func isValid(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	value := node.Val

	if lower != nil {
		if value <= *lower {
			return false
		}
	}

	if upper != nil {
		if value >= *upper {
			return false
		}
	}

	return isValid(node.Left, lower, &value) && isValid(node.Right, &value, upper)
}
