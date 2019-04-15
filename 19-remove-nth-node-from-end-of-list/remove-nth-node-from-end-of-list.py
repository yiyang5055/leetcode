# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        tmp_head = ListNode(None)
        tmp_head.next = head
        
        fast = tmp_head
        slow = tmp_head
        while n:
            fast = fast.next
            n = n - 1
        
        while fast.next:
            fast = fast.next
            slow = slow.next
        
        tmp = slow.next
        slow.next = slow.next.next
        del tmp
        return tmp_head.next
        