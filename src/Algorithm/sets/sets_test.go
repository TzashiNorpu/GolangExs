package sets

import "testing"

func TestSets(t *testing.T) {
	nums := []int{1, 2, 3}
	generateAllSets(nums, 0, make([]int, 0))
}
