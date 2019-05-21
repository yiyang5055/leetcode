/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (56.04%)
 * Total Accepted:    442.8K
 * Total Submissions: 789.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 *
 */

//  Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type ListNode struct {
	Val  *TreeNode
	Next *ListNode
}
type Stack struct {
	Head *ListNode
}

func (s *Stack) Push(root *TreeNode) {
	cur := &ListNode{
		Val:  root,
		Next: s.Head,
	}

	s.Head = cur
}
func (s *Stack) Pop() *TreeNode {
	cur := s.Head
	s.Head = s.Head.Next
	cur.Next = nil
	return cur.Val
}
func (s *Stack) Empty() bool {
	return s.Head == nil
}

func inorderRecursiveTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderRecursiveTraversal(root.Left, res)
	*res = append(*res, root.Val)
	inorderRecursiveTraversal(root.Right, res)
}

func inorderIterativeTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := &Stack{}
	res := []int{}
	cur := root
	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop()
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

func inorderTraversal(root *TreeNode) []int {
	return inorderIterativeTraversal(root)
}
