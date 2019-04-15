# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def binaryTreePaths(self, root):
        """
        :type root: TreeNode
        :rtype: List[str]
        """
        cur_path = list()
        result = list()
        self._binaryTreePaths(root, cur_path, result)
        return result
        
    
    def _binaryTreePaths(self, root, cur_path, result):
        if not root: return None
        
        cur_path.append(str(root.val))
        if not root.left and not root.right:
            result.append("->".join(cur_path))
        self._binaryTreePaths(root.left, cur_path, result)
        self._binaryTreePaths(root.right, cur_path, result)
        cur_path.pop()
        