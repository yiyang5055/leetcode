/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	walker, runner := head, head
	for runner.Next != nil && runner.Next.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}

	return false
}
