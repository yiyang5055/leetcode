# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def convertToBST(self, nums , start, end):
        if start > end:
            return None
        
        mid = start + (end - start) / 2
        parent = TreeNode(nums[mid])
        parent.left = self.convertToBST(nums, start, mid - 1)
        parent.right = self.convertToBST(nums, mid + 1, end)
        
        return parent
    
    def sortedArrayToBST(self, nums):
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        nums_len = len(nums)
        tree_root = self.convertToBST(nums, 0, nums_len - 1)
        return tree_root
        
        