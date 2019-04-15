# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def preorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        res = list()
        path_stack = list()
        node = root
        
        while node or len(path_stack) > 0:
            while node:
                res.append(node.val)
                path_stack.append(node)
                node = node.left
                
            if len(path_stack):
                node = path_stack.pop()
                node = node.right
        return res
        