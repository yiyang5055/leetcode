class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        last = len(nums)
        nums.sort()
        min_gap = float("inf")
        res = 0
        for i in xrange(last - 2):
            j = i + 1
            k = last - 1
            while j < k:
                sum = nums[i] + nums[j] + nums[k]
                gap = abs(target - sum)
                if gap < min_gap:
                    res = sum
                    min_gap = gap
                
                if sum < target:
                    j += 1
                elif sum > target:
                    k -= 1
                else:
                    return target

        return res
        