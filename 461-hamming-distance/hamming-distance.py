class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        value = x ^ y
        num = 0
        while value:
            value = value & (value - 1)
            num += 1
        return num