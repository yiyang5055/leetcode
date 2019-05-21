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

func validBST(root *TreeNode, low, upper float64) bool {
	if root == nil {
		return true
	}

	if float64(root.Val) <= low || float64(root.Val) >= upper {
		return false
	}

	return validBST(root.Left, low, float64(root.Val)) && validBST(root.Right, float64(root.Val), upper)
}

func validBSTIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := &Stack{}
	pre := (*TreeNode)(nil)
	cur := root
	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop()
		if pre != nil && cur.Val <= pre.Val {
			return false
		}
		pre = cur
		cur = cur.Right
	}
	return true
}
func isValidBST(root *TreeNode) bool {
	return validBSTIterative(root)
}
