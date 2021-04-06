package lt80

func removeDuplicates(nums []int) int {
	return process(nums, 2)
}

func process(nums []int, i int) int {
	u := 0
	for _, v := range nums {
		if u < i || nums[u-i] != v {
			//nums[u++] = v
			nums[u] = v
			u += 1
		}
	}
	return u
}
