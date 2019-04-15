# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        result, _ = self._isBalanced(root)
        return result
    
    def _isBalanced(self, root):
        if not root:
            return True, 0
        
        result_left, height_left = self._isBalanced(root.left)
        result_right, height_right = self._isBalanced(root.right)
        
        if result_left and result_right and abs(height_left - height_right) <= 1:
            return True, max(height_left, height_right) + 1
        else:
            return False, max(height_left, height_right) + 1