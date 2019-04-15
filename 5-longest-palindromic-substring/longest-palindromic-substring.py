class Solution(object):
    def palindromic(self, s, front, after):
        str_len = len(s)
        front_index = front
        after_index = after

        while front_index >= 0 and after_index <= str_len - 1 and\
                s[front_index] == s[after_index]:
            front_index -= 1
            after_index += 1

        return s[front_index + 1:after_index]

    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        str_len = len(s)
        if str_len == 0:
            return None

        longest_substr = s[0]
        for index in xrange(str_len - 1):
            tmp_str = self.palindromic(s, index, index)
            if len(tmp_str) > len(longest_substr):
                longest_substr = tmp_str
            tmp_str = self.palindromic(s, index, index + 1)
            if len(tmp_str) > len(longest_substr):
                longest_substr = tmp_str

        return longest_substr
        