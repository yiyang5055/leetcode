/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (52.14%)
 * Total Accepted:    378.5K
 * Total Submissions: 721.1K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
func majorityElement(nums []int) int {
	res := 0
	n := 0
	for i := 0; i < len(nums); i++ {
		if n == 0 {
			res = nums[i]
			n = 1
		} else {
			if res == nums[i] {
				n++
			} else {
				n--
			}
		}
	}
	return res
}
