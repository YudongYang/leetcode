/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		inList = make([]int, numCourses)
		l      = make([]int, 0, numCourses)
		result = make([]int, 0, numCourses)
	)

	for _, e := range prerequisites {
		edges[e[1]] = append(edges[e[1]], e[0])
		inList[e[0]]++
	}

	for i, e := range inList {
		if e == 0 {
			l = append(l, i)
		}
	}

	for len(l) > 0 {
		head := l[0]
		l = l[1:]
		result = append(result, head)
		for _, e := range edges[head] {
			inList[e]--
			if inList[e] == 0 {
				l = append(l, e)
			}
		}
	}

	if len(result) != numCourses {
		return nil
	}

	return result
}
// @lc code=end

