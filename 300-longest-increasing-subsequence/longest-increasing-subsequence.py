class Solution(object):
    def lengthOfLIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
            
        length_list = [1] * len(nums)
        for i in xrange(len(nums)):
            for j in xrange(0, i + 1):
                if nums[i] > nums[j] and length_list[j] + 1 > length_list[i]:
                    length_list[i] = length_list[j] + 1

        return max(length_list)