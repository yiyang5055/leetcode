# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def sumOfLeftLeaves(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        return self._sumOfLeftLeaves(root, False)
    
    def _sumOfLeftLeaves(self, root, is_left):
        if not root:
            return 0
        
        if not root.left and not root.right and is_left:
            return root.val
        
        return self._sumOfLeftLeaves(root.left, True) + \
               self._sumOfLeftLeaves(root.right, False)
        