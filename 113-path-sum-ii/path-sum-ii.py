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
        :rtype: List[List[int]]
        """
        result = list()
        cur_path = list()
        self._pathSum(root, cur_path, sum, result)
        return result
        
    def _pathSum(self, root, cur_path, sum, result):
        if not root: return
        
        cur_path.append(root.val)
        if root.val == sum and not root.left and not root.right:
            result.append(cur_path[:])
        
        self._pathSum(root.left, cur_path, sum - root.val, result)
        self._pathSum(root.right, cur_path, sum - root.val, result)
        cur_path.pop()
        