/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (33.32%)
 * Total Accepted:    287.1K
 * Total Submissions: 861.7K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 */
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) < 1 {
		return res
	}

	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if target == nums[low] {
		res[0] = low
	} else {
		return res
	}

	low, high = res[0], len(nums)-1
	for low < high {
		mid := (low+high)/2 + 1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}

	if target == nums[low] {
		res[1] = high
	} else {
		return res
	}

	return res
}
