# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def kthSmallest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        
        stack = list()
        p = root
        
        while(len(stack) > 0 or p):
            if p:
                stack.append(p)
                p = p.left
            else:
                p = stack.pop()
                k = k - 1
                if k == 0:
                    return p.val
                p = p.right
                
        return -1
        
        