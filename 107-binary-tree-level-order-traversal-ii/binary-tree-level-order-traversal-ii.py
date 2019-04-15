# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        res = list()
        current, next = list(), list()
        
        if not root:
            return res  
        current.append(root)
        
        while len(current) > 0:
            level = list()
            while len(current) > 0:
                node = current.pop(0)
                level.append(node.val)
                if node.left:
                    next.append(node.left)
                if node.right:
                    next.append(node.right)
            res.append(level)
            current, next = next, current
            
        res.reverse()
        return res 
        