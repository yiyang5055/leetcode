/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	p := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    l1 := mergeKLists(lists[0:len(lists)/2])
    l2 := mergeKLists(lists[len(lists)/2:])

    return mergeTwoLists(l1, l2)
    
}