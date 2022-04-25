/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, e := range nums {
		if j, ok := m[target - e]; ok {
			return []int{j, i}
		}
		m[e] = i
	}
	return []int{}
}
// @lc code=end

