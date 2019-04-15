# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    
    def convertToBST(self, lists , start, end):
        if start > end:
            return None
        
        mid = start + (end - start) / 2
        parent = TreeNode(lists[mid].val)
        parent.left = self.convertToBST(lists, start, mid - 1)
        parent.right = self.convertToBST(lists, mid + 1, end)
        
        return parent
    
    def sortedListToBST(self, head):
        """
        :type head: ListNode
        :rtype: TreeNode
        """
        link_list = list()
        link_len = 0
        tmp_head = head
        
        while tmp_head:
            link_list.append(tmp_head)
            tmp_head = tmp_head.next
            link_len += 1
        return self.convertToBST(link_list, 0, link_len - 1)
        
        
        