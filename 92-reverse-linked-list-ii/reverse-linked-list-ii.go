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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n == 1 {
		return head
	}

	res := &ListNode{
		Next: head,
	}

	pre := res
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	last := pre
	for i := 0; i <= n-m; i++ {
		last = last.Next
	}
	tail := last.Next
	last.Next = nil

	reverse := reverseList(pre.Next)
	pre.Next = reverse
	for pre.Next != nil {
		pre = pre.Next
	}
	pre.Next = tail

	return res.Next
}
