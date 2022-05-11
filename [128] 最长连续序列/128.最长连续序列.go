/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, e := range nums {
		m[e] = struct{}{}
	}

	max := 0

	for _, e := range nums {
		if _, ok := m[e-1]; !ok {
			num := e
			n := 1
			for {
				if _, ok := m[num+1]; !ok {
					break
				}
				num = num + 1
				n++
			}
			if n > max {
				max = n
			}
		}
	}

	return max
}
// @lc code=end

