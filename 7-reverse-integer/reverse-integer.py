class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        res = 0
        tmp = abs(x)
        while tmp >= 1:
            res = (res * 10 + tmp % 10)
            tmp = tmp / 10

        res = res if res <= 0x7fffffff else 0
        return -res if x < 0 else res
        