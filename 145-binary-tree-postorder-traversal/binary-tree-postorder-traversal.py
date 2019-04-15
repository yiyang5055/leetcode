# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class TraversalNode(object):
    def __init__(self, node):
        self.is_first = True
        self.node = node

class Solution(object):
    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        res = list()
        path_stack = list()
        
        while root or len(path_stack) > 0:
            while root:
                traversal_node = TraversalNode(root)
                path_stack.append(traversal_node)
                root = root.left
                
            if len(path_stack) > 0:
                traversal_node = path_stack.pop()
                if traversal_node.is_first:
                    traversal_node.is_first = False
                    path_stack.append(traversal_node)
                    root = traversal_node.node.right
                else:
                    res.append(traversal_node.node.val)
                    root = None
        
        return res
        