/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (25.25%)
 * Total Accepted:    186.5K
 * Total Submissions: 735.3K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 *
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 *
 *
 * Example 2:
 *
 *
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 *
 *
 */
type ListNode struct {
	Next *ListNode
	Val  int
}

type Stack struct {
	Head *ListNode
}

func (stack *Stack) Pop() {
	if stack.Empty() {
		return
	}
	h := stack.Head
	stack.Head = stack.Head.Next
	h.Next = nil
}

func (stack *Stack) Push(val int) {
	h := &ListNode{
		Val:  val,
		Next: stack.Head,
	}
	stack.Head = h
}

func (stack *Stack) Peek() int {
	if stack.Empty() {
		return -1
	}

	return stack.Head.Val
}

func (stack *Stack) Empty() bool {
	return stack.Head == nil
}

func longestValidParentheses(s string) int {
	stack := &Stack{}
	validLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || stack.Empty() {
			stack.Push(i)
		} else {
			peek := stack.Peek()
			if s[peek] == '(' {
				stack.Pop()
				if i-stack.Peek() > validLen {
					validLen = i - stack.Peek()
				}
			} else {
				stack.Push(i)
			}
		}
	}

	return validLen
}
