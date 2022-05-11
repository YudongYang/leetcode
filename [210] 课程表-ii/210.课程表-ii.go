/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges      = make([][]int, numCourses)
		statusList = make([]int, numCourses)
		result     = make([]int, 0, numCourses)
		cycle      bool
		dfs        func(i int)
	)

	dfs = func(i int) {
		statusList[i] = 1
		for _, e := range edges[i] {
			if statusList[e] == 0 {
				dfs(e)
				if cycle {
					return
				}
			} else if statusList[e] == 1 {
				cycle = true
				return
			}
		}
		statusList[i] = 2
		result = append(result, i)
	}

	for _, e := range prerequisites {
		edges[e[1]] = append(edges[e[1]], e[0])
	}

	for i := 0; i < numCourses && !cycle; i++ {
		if statusList[i] == 0 {
			dfs(i)
		}
	}

	if cycle {
		return nil
	}

	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}
// @lc code=end

