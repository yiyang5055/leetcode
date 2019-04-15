# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def hasPathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: bool
        """
        
        if not root: return False
        if root.val == sum and not root.left and not root.right: 
            return True
        
        left_result = self.hasPathSum(root.left, sum - root.val)
        right_result = self.hasPathSum(root.right, sum - root.val)
        return left_result or right_result
        