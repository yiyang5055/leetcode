# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeTwoLists(self, list1, list2):

        head = ListNode(None)
        p = head

        while list1 and list2:
            if list1.val <= list2.val:
                p.next = list1
                list1 = list1.next
            else:
                p.next = list2
                list2 = list2.next
            p = p.next

        p.next = list1 if list1 else list2

        return head.next

    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        if not lists:
            return None
            
        if len(lists) == 1:
            return lists[0]

        lists_len = len(lists)
        res1 = self.mergeKLists(lists[0:lists_len / 2])
        res2 = self.mergeKLists(lists[lists_len / 2:])
        return self.mergeTwoLists(res1, res2)
        