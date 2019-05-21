/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 *
 * https://leetcode.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (40.72%)
 * Total Accepted:    311K
 * Total Submissions: 763.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, determine if it is height-balanced.
 *
 * For this problem, a height-balanced binary tree is defined as:
 *
 *
 * a binary tree in which the depth of the two subtrees of every node never
 * differ by more than 1.
 *
 *
 * Example 1:
 *
 * Given the following tree [3,9,20,null,null,15,7]:
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * Return true.
 *
 * Example 2:
 *
 * Given the following tree [1,2,2,3,3,null,null,4,4]:
 *
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * Return false.
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := dfsHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := dfsHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}
func isBalanced(root *TreeNode) bool {
	return dfsHeight(root) != -1
}
