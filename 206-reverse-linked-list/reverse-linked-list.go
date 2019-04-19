/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (53.85%)
 * Total Accepted:    557.1K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you
 * implement both?
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//2->1->3->4->5

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	resHead := head
	tail := head
	for tail.Next != nil {
		currentNode := resHead
		resHead = tail.Next
		tail.Next = resHead.Next
		resHead.Next = currentNode

	}

	return resHead
}
