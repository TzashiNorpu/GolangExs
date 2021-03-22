package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	generateAllSets(nums, 0, make([]int, 0))
}

func generateAllSets(nums []int, level int, res []int) {
	if level == len(nums) {
		fmt.Println(res)
		return
	}
	generateAllSets(nums, level+1, res)
	generateAllSets(nums, level+1, append(res, nums[level]))
}
