class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        num_map = {
            "I": 1, 
            "V": 5, 
            "X": 10, 
            "L": 50, 
            "C":100, 
            "D": 500, 
            "M": 1000
        }
        
        res = num_map[s[0]]
        for i in xrange(1, len(s)):
            if num_map[s[i]] > num_map[s[i - 1]]:
                res += num_map[s[i]] - 2 * num_map[s[i - 1]]
            else:
                res += num_map[s[i]]
        return res