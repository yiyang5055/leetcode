class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for index in xrange(len(nums) + 1):
            tmp = target - nums[index]
            sub_nums = nums[index + 1:]
            if tmp in sub_nums:
                return [index, sub_nums.index(tmp) + index + 1]