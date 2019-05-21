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
func (s *Stack) Push(val *TreeNode) {
	cur := &ListNode{
		Val:  val,
		Next: s.Head,
	}

	s.Head = cur
}
func (s *Stack) Empty() bool {
	return s.Head == nil
}

func preorderRecursiveTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	preorderRecursiveTraversal(root.Left, res)
	preorderRecursiveTraversal(root.Right, res)
}

func preorderIterativeTraversal(root *TreeNode) []int {
	res := []int{}
	stack := &Stack{}
	cur := root
	for cur != nil || !stack.Empty() {
		for cur != nil {
			res = append(res, cur.Val)
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop()
		cur = cur.Right
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	return preorderIterativeTraversal(root)
}
