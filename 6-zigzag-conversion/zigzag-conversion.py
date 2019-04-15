class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1 or len(s) == 0:
            return s

        res = list()
        str_len = len(s)
        row_index = 0
        while row_index < str_len and row_index < numRows:
            s_index = row_index
            res.append(s[s_index])
            column_index = 1
            while s_index < str_len:
                if row_index == 0 or row_index == numRows - 1:
                    s_index += 2 * numRows - 2
                elif column_index & 0x1:
                    s_index += 2 * (numRows - 1 - row_index)
                else:
                    s_index += 2 * row_index
                
                if s_index < str_len:
                    res.append(s[s_index])

                column_index += 1
            row_index += 1

        return ''.join(res)
        