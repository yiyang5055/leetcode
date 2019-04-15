# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        res = list()
        path_stack = list()
        
        while root or len(path_stack) > 0:
            while root:
                path_stack.append(root)
                root = root.left
                
            if len(path_stack):
                root = path_stack.pop()
                res.append(root.val)
                root = root.right
        return res
        