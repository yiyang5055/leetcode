class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        parentheses_map = {
            ")": "(",
            "}": "{",
            "]": "["
        }
        
        parentheses_stack = list()
        
        for value in s:
            if value in "({[" or len(parentheses_stack) == 0:
                parentheses_stack.append(value)
                continue
            
            if parentheses_stack[-1] == parentheses_map[value]:
                parentheses_stack.pop()
            else:
               parentheses_stack.append(value)
                    
        return True if len(parentheses_stack) == 0 else False 

                    