class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        sub_set = set()
        str_length = len(s)
        i, j = 0, 0
        count = 0
        while i < str_length and j < str_length:
            if s[j] not in sub_set:
                sub_set.add(s[j])
                j += 1
                count = max(count, j - i)
            else:
                sub_set.remove(s[i])
                i += 1

        return count