package main

import "fmt"

func dsp(nums []int, x, y int, result *[][]int) {
	// fmt.Println(x, y)
	if x >= len(nums) || y >= len(nums) {
		return
	}
	newNums := nums
	fmt.Println(newNums)
	newNums[x], newNums[y] = newNums[y], newNums[x]
	for j := y; j < len(nums); j++ {
		dsp(newNums, x, y+1, result)
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		dsp(nums, i, 0, &result)
	}
	return result
}

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(permute(a))
}
