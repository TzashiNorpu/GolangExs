package lt80

import (
	"fmt"
	"testing"
)

func TestLt80(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 3}
	s := removeDuplicates(nums)
	fmt.Println(s)
}
