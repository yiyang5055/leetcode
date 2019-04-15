# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: int
        """
        if not root: return 0
        return self._dfs(root, sum) + self.pathSum(root.left, sum) + self.pathSum(root.right, sum)
    
    def _dfs(self, root, sum):
        result = 0
        if not root: return result
    
        if root.val == sum:
            result += 1
        
        result += self._dfs(root.left, sum - root.val)
        result += self._dfs(root.right, sum - root.val)
        
        return result
        
        