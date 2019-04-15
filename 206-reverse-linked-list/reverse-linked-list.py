# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        reverse_head = None
        
        while head:
            current_node = head
            head = head.next
            current_node.next = reverse_head
            reverse_head = current_node
        
        return reverse_head