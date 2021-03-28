package lt83

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt83(t *testing.T) {
	ints := []int{1, 1, 2, 3, 3, 3, 3}
	l := list.GenLinkedListByInts(ints)
	deleteDuplicates(l.Node)
	l.Traverse()
}
