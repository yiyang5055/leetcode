# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def _isValidBST(self, root, lower, upper):
        if not root:
            return True
            
        return root.val > lower and root.val < upper \
               and self._isValidBST(root.left, lower, root.val) \
               and self._isValidBST(root.right, root.val, upper)
               
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self._isValidBST(root, float("-inf"), float("inf"))
        