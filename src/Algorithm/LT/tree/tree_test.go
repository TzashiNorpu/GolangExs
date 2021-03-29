package tree

import "testing"

func TestTree(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 7, 8, 9}
	l := GenTreesByArray(ints)
	inOrder(l)
}
