package lc0173

import (
	"leetcode-solutions-golang-zea7ot/src/leetcode/util/ds"
)

// TreeNode is the binary tree node
type TreeNode = ds.TreeNode

// BSTIterator will be instantiated and called as such:
// obj := Constructor(root);
// param_1 := obj.Next();
// param_2 := obj.HasNext();
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor is
func Constructor(root *TreeNode) BSTIterator {
	var iter BSTIterator
	iter.PushLeft(root)
	return iter
}

// Next returns the next tree node value in inorder sequence
func (iter *BSTIterator) Next() int {
	top := iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]
	iter.PushLeft(top.Right)
	return top.Val
}

// HasNext returns whether the binary tree has completely been explored
func (iter *BSTIterator) HasNext() bool {
	return len(iter.stack) > 0
}

// PushLeft pushes the all left tree nodes into the stack
func (iter *BSTIterator) PushLeft(node *TreeNode) {
	for node != nil {
		iter.stack = append(iter.stack, node)
		node = node.Left
	}
}
