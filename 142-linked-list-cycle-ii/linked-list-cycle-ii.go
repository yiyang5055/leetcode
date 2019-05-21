/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	walker, runner := head, head
	for runner.Next != nil && runner.Next.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			for cur := head; cur != walker; cur = cur.Next {
				walker = walker.Next
			}
			return walker
		}
	}
	return nil
}
