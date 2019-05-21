/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (27.01%)
 * Total Accepted:    546.5K
 * Total Submissions: 2M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */
func palindrom(s string, begin, end int) string {
	for begin >= 0 && end < len(s) && s[begin] == s[end] {
		begin--
		end++
	}
	begin++
	return s[begin:end]
}
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	res := ""
	for i := 0; i < len(s)-1; i++ {
		p := palindrom(s, i, i)
		if len(p) >= len(res) {
			res = p
		}

		p = palindrom(s, i, i+1)
		if len(p) >= len(res) {
			res = p
		}
	}
	return res
}
