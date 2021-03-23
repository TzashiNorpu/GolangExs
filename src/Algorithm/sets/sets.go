package sets

import "fmt"

func generateAllSets(nums []int, level int, res []int) {
	if level == len(nums) {
		fmt.Println(res)
		return
	}
	generateAllSets(nums, level+1, res)
	generateAllSets(nums, level+1, append(res, nums[level]))
}
