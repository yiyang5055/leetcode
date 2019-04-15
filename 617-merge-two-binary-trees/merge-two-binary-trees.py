# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def mergeTrees(self, t1, t2):
        """
        :type t1: TreeNode
        :type t2: TreeNode
        :rtype: TreeNode
        """
        result = None
        if t1 and t2:
            result = TreeNode(t1.val + t2.val)
            result.left = self.mergeTrees(t1.left, t2.left)
            result.right = self.mergeTrees(t1.right, t2.right)
        elif t1:
            result = TreeNode(t1.val)
            result.left = self.mergeTrees(t1.left, None)
            result.right = self.mergeTrees(t1.right, None)
        elif t2:
            result = TreeNode(t2.val)
            result.left = self.mergeTrees(None, t2.left)
            result.right = self.mergeTrees(None, t2.right)

        return result
            
        
        