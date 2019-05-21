/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 *
 * https://leetcode.com/problems/sum-of-two-integers/description/
 *
 * algorithms
 * Easy (51.00%)
 * Total Accepted:    133.3K
 * Total Submissions: 261.6K
 * Testcase Example:  '1\n2'
 *
 * Calculate the sum of two integers a and b, but you are not allowed to use
 * the operator + and -.
 *
 *
 * Example 1:
 *
 *
 * Input: a = 1, b = 2
 * Output: 3
 *
 *
 *
 * Example 2:
 *
 *
 * Input: a = -2, b = 3
 * Output: 1
 *
 *
 *
 *
 */
func getSum(a int, b int) int {
	for b != 0 {
		c := a ^ b
		b = a & b << 1
		a = c
	}
	return a
}
