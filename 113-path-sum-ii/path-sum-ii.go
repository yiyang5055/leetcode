/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (40.18%)
 * Total Accepted:    224.8K
 * Total Submissions: 559K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's
 * sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 * ⁠     5
 * ⁠    / \
 * ⁠   4   8
 * ⁠  /   / \
 * ⁠ 11  13  4
 * ⁠/  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
 * ]
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSumII(root *TreeNode, sum int, res [][]int, current []int) [][]int {
	if root == nil {
		return res
	}

	current = append(current, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		res = append(res, append([]int(nil), current...))
		current = current[:len(current)-1]
	}
	res = pathSumII(root.Left, sum-root.Val, res, current)
	res = pathSumII(root.Right, sum-root.Val, res, current)
	return res
}
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var current []int
	return pathSumII(root, sum, res, current)
}
