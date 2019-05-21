/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for cur := head; cur != nil && cur.Next != nil; {
		next := cur.Next
		cur.Next = cur.Next.Next
		next.Next = cur
		pre.Next = next
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}
