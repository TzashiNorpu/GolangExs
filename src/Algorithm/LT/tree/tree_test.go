package tree

import "testing"

func TestTree(t *testing.T) {
	ints := []int{1, 2, 3}
	l := GenTreesByArray(ints)
	inOrder(l)
}
