/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 *
 * https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (63.62%)
 * Total Accepted:    123.7K
 * Total Submissions: 193.4K
 * Testcase Example:  '"Let\'s take LeetCode contest"'
 *
 * Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be
 * any extra space in the string.
 *
 */
func reverseWords(s string) string {
	sb := []byte(s)
	for h := 0; h < len(sb)-1; {
		b := h
		for b < len(sb) && sb[b] != ' ' {
			b++
		}

		i, j := h, b-1
		for i <= j {
			sb[i], sb[j] = sb[j], sb[i]
			i++
			j--
		}
		h = b + 1
	}

	return string(sb)
}
