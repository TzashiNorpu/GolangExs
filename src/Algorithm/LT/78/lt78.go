package lt78

func subsets(nums []int) [][]int {
	var res []int
	var r [][]int
	generateAllSets(nums, 0, res, &r)
	return r
}

func generateAllSets(nums []int, level int, res []int, r *[][]int) {
	if level == len(nums) {
		//fmt.Println(res)
		*r = append(*r, res)
		return
	}

	sel := append(res, nums[level])
	generateAllSets(nums, level+1, sel, r)
	generateAllSets(nums, level+1, res, r)
}

//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
