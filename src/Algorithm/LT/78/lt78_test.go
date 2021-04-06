package lt78

import (
	"fmt"
	"testing"
)

func TestLt78(t *testing.T) {
	ints := []int{9, 0, 3, 5, 7}
	res := subsets(ints)
	fmt.Println(res)
}
