class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        num = 0
        
        while n:
            n = n & (n - 1)
            num += 1
        return num
                
        