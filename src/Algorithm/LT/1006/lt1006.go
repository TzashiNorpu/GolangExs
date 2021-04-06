package lt1006

func clumsy(N int) int {
	//12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1
	ops := []byte{'*', '/', '+', '-'}
	nums := []int{}
	nums = append(nums, N)
	for i := N - 1; i >= 1; i-- {
		op := ops[0]
		ops = ops[1:]
		if op == '*' {
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums = append(nums, num*i)
			ops = append(ops, '*')
		} else if op == '/' {
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums = append(nums, num/i)
			ops = append(ops, '/')
		} else if op == '-' {
			nums = append(nums, -i)
			ops = append(ops, '-')
		} else {
			nums = append(nums, i)
			ops = append(ops, '+')
		}
	}
	res := 0
	for _, v := range nums {
		res = res + v
	}
	return res
}
