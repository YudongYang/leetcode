/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max = new(int)
	*max = math.MinInt

	maxNode(root, max)

	return *max
}

func maxNode(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}

	val := node.Val
	if node.Left == nil && node.Right == nil {
		*max = maxArray(val, *max)
	} else {
		maxLeft := maxNode(node.Left, max)
		maxRight := maxNode(node.Right, max)

		node.Val = maxArray(val, val+maxLeft, val+maxRight)

		*max = maxArray(node.Val, val+maxLeft+maxRight, *max)
	}
	return node.Val
}

func maxArray(l ...int) (max int) {
	max = math.MinInt
	for _, e := range l {
		if e > max {
			max = e
		}
	}
	return max
}
// @lc code=end

