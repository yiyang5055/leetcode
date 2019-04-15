class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        res = list()
        for i in xrange(num + 1):
            j = i
            num = 0
            while j:
                j = j & (j - 1)
                num += 1
            res.append(num)
        return res