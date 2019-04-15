# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """
        if not preorder: return None
        if not inorder: return None
        
        root = TreeNode(preorder[0])

        if preorder[0] not in inorder: return None
        root_index = inorder.index(preorder[0])
        
        root.left = self.buildTree(preorder[1:], inorder[:root_index])
        root.right = self.buildTree(preorder[root_index + 1:], inorder[root_index + 1:])
        
        return root