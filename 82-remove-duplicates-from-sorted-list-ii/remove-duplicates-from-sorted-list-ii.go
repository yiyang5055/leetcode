/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Val:  head.Val - 1,
		Next: head,
	}
	pre := dummy
	for pre.Next != nil {
		cur := pre.Next
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if cur != pre.Next {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}
