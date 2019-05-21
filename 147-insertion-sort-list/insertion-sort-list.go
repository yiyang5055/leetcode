/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		cur := head
		for h := dummy; h != nil; h = h.Next {
			if h.Next == nil || head.Val < h.Next.Val {
				head = head.Next
				cur.Next = h.Next
				h.Next = cur
				break
			}
		}
	}
	return dummy.Next
}
