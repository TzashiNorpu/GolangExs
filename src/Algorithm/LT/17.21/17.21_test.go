package trap

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	ints := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	s := trap2(ints)
	fmt.Println(s)
}
