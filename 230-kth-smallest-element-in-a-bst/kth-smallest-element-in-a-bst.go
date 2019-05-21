/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode struct {
	Val  *TreeNode
	Next *ListNode
}

type Stack struct {
	Head *ListNode
}

func (s *Stack) Pop() *TreeNode {
	cur := s.Head
	s.Head = s.Head.Next
	cur.Next = nil
	return cur.Val
}
func (s *Stack) Push(n *TreeNode) {
	cur := &ListNode{
		Val:  n,
		Next: s.Head,
	}

	s.Head = cur
}
func (s *Stack) Empty() bool {
	return s.Head == nil
}

func kthSmallest(root *TreeNode, k int) int {
	stack := &Stack{}
	cur := root
	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop()
		k -= 1
		if k == 0 {
			break
		}
		cur = cur.Right
	}
	return cur.Val
}
