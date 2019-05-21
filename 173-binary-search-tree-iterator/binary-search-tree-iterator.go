/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 *
 * https://leetcode.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (48.04%)
 * Total Accepted:    202.4K
 * Total Submissions: 417.7K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n[[[7,3,15,null,null,9,20]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]'
 *
 * Implement an iterator over a binary search tree (BST). Your iterator will be
 * initialized with the root node of a BST.
 *
 * Calling next() will return the next smallest number in the BST.
 *
 *
 *
 *
 *
 *
 * Example:
 *
 *
 *
 *
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // return 3
 * iterator.next();    // return 7
 * iterator.hasNext(); // return true
 * iterator.next();    // return 9
 * iterator.hasNext(); // return true
 * iterator.next();    // return 15
 * iterator.hasNext(); // return true
 * iterator.next();    // return 20
 * iterator.hasNext(); // return false
 *
 *
 *
 *
 * Note:
 *
 *
 * next() and hasNext() should run in average O(1) time and uses O(h) memory,
 * where h is the height of the tree.
 * You may assume that next() call will always be valid, that is, there will be
 * at least a next smallest number in the BST when next() is called.
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type StackNode struct {
	Val  *TreeNode
	Next *StackNode
}

type Stack struct {
	Head *StackNode
}

func (s *Stack) Push(node *TreeNode) {
	n := &StackNode{
		Val:  node,
		Next: s.Head,
	}

	s.Head = n
}

func (s *Stack) Pop() *TreeNode {
	if s.Empty() {
		return nil
	}
	t := s.Head
	s.Head = s.Head.Next

	return t.Val
}

func (s *Stack) Empty() bool {
	return s.Head == nil
}

type BSTIterator struct {
	stack *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{
		stack: &Stack{},
	}

	bst.push(root)
	return bst
}

func (this *BSTIterator) Next() int {
	node := this.stack.Pop()
	this.push(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return !this.stack.Empty()
}

func (this *BSTIterator) push(root *TreeNode) {
	for node := root; node != nil; node = node.Left {
		this.stack.Push(node)
	}
}
