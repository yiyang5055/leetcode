/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	walker, runnner := head, head.Next
	for runnner.Next != nil {
		runnner = runnner.Next
		walker = walker.Next
		if runnner.Next == nil {
			break
		}
		runnner = runnner.Next
	}
    n := walker.Next
	walker.Next = nil

	return head, n
}
func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else if l2 == nil {
		tail.Next = l1
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l1, l2 := splitList(head)
	head1 := sortList(l1)
	head2 := sortList(l2)
	return mergeList(head1, head2)
}