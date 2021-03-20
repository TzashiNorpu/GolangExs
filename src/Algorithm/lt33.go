package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if nums[l] <= nums[mid] {
			if target <= nums[mid] && target >= nums[l] {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[r] {
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	t := 7
	s := search(a, t)
	fmt.Println(s)
	//for i, _ := range a {
	//	fmt.Println(i)
	//}

}
