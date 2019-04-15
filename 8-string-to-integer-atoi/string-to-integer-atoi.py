class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        res = 0
        digit_str = str.strip()
        if len(digit_str) == 0:
            return 0

        digit_substr = digit_str
        if digit_str[0] in '+-':
            digit_substr = digit_str[1:]

        for item in digit_substr:
            if item not in '0123456789':
                break
            if res > 0x7fffffff / 10 \
                    or (res == 0x7fffffff / 10 and ord(item) - ord('0') > 0x7fffffff % 10):
                return -0x80000000 if digit_str[0] == "-" else 0x7fffffff

            res = res * 10 + (ord(item) - ord('0'))

        return -res if digit_str[0] == "-" else res