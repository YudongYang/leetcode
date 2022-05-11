/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges      = make([][]int, numCourses)
		statusList = make([]int, numCourses)
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
	}

	for _, e := range prerequisites {
		edges[e[1]] = append(edges[e[1]], e[0])
	}

	for i := 0; i < numCourses && !cycle; i++ {
		if statusList[i] == 0 {
			dfs(i)
		}
	}

	return !cycle
}
// @lc code=end

